package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	bffpb "ping-pong/bff"
	exclamationpb "ping-pong/exclamation"
	pongpb "ping-pong/pong"
)

type BFFServer struct {
	bffpb.UnimplementedBFFServiceServer
}

func (s *BFFServer) ForwardPing(ctx context.Context, req *bffpb.BFFRequest) (*bffpb.BFFResponse, error) {
	// Pongサービスにリクエスト
	pongResponse, err := callPongService(req.Message)
	if err != nil {
		return nil, err
	}

	// Exclamationサービスにリクエスト
	exclamationResponse, err := callExclamationService(req.Message)
	if err != nil {
		return nil, err
	}

	// レスポンスを統合
	combinedResponse := pongResponse + exclamationResponse
	return &bffpb.BFFResponse{Response: combinedResponse}, nil
}

// callPongServiceはPongサービスにgRPCリクエストを送信します
func callPongService(message string) (string, error) {
	// gRPC接続を確立
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials())) // Pongサービスのアドレス
	if err != nil {
		log.Printf("Failed to connect to Pong Service: %v", err)
		return "", err
	}
	defer conn.Close()

	client := pongpb.NewPongServiceClient(conn)
	req := &pongpb.PongRequest{Message: message}
	res, err := client.GetPong(context.Background(), req)
	if err != nil {
		log.Printf("Error calling Pong Service: %v", err)
		return "", err
	}

	return res.Response, nil
}

// callExclamationServiceはExclamationサービスにgRPCリクエストを送信します
func callExclamationService(message string) (string, error) {
	// gRPC接続を確立
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials())) // Exclamationサービスのアドレス
	if err != nil {
		log.Printf("Failed to connect to Exclamation Service: %v", err)
		return "", err
	}
	defer conn.Close()

	client := exclamationpb.NewExclamationServiceClient(conn)
	req := &exclamationpb.ExclamationRequest{Message: message}
	res, err := client.GetExclamation(context.Background(), req)
	if err != nil {
		log.Printf("Error calling Exclamation Service: %v", err)
		return "", err
	}

	return res.Response, nil
}

func main() {
	// gRPCサーバーを設定
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	bffpb.RegisterBFFServiceServer(s, &BFFServer{})

	log.Println("BFF running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
