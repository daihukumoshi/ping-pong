package main

import (
	"context"
	"log"
	"net"

	pongpb "ping-pong/pong"

	"google.golang.org/grpc"
)

type server struct {
	pongpb.UnimplementedPongServiceServer
}

func (s *server) GetPong(ctx context.Context, req *pongpb.PongRequest) (*pongpb.PongResponse, error) {
	return &pongpb.PongResponse{Response: "Pong"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pongpb.RegisterPongServiceServer(s, &server{})

	log.Println("Pong Service running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
