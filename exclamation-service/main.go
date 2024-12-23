package main

import (
	"context"
	"log"
	"net"

	exclamationpb "ping-pong/exclamation"

	"google.golang.org/grpc"
)

type server struct {
	exclamationpb.UnimplementedExclamationServiceServer
}

func (s *server) GetExclamation(ctx context.Context, req *exclamationpb.ExclamationRequest) (*exclamationpb.ExclamationResponse, error) {
	return &exclamationpb.ExclamationResponse{Response: "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	exclamationpb.RegisterExclamationServiceServer(s, &server{})

	log.Println("Exclamation Service running on port 50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
