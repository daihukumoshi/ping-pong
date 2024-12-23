package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	bffpb "ping-pong/bff"
)

type PingRequest struct {
	Message string `json:"message"`
}

type PingResponse struct {
	Response string `json:"response"`
}

func main() {
	http.HandleFunc("/ping", handlePing) // HTTPエンドポイントを設定

	log.Println("API Gateway running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	// HTTPリクエストのパース
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req PingRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// gRPCクライアントでBFFに転送
	response, err := forwardToBFF(req.Message)
	if err != nil {
		http.Error(w, "Failed to forward request to BFF", http.StatusInternalServerError)
		return
	}

	// HTTPレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PingResponse{Response: response})
}

func forwardToBFF(message string) (string, error) {
	// gRPC接続を確立
	conn, err := grpc.NewClient("dns:///localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to BFF: %v", err)
		return "", err
	}
	defer conn.Close()

	client := bffpb.NewBFFServiceClient(conn)

	// gRPCリクエストを作成して送信
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &bffpb.BFFRequest{Message: message}
	res, err := client.ForwardPing(ctx, req)
	if err != nil {
		log.Printf("Failed to call BFF: %v", err)
		return "", err
	}

	return res.Response, nil
}
