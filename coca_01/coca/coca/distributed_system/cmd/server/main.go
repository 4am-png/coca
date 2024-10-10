package main

import (
	"fmt"
	"net"
	"google.golang.org/grpc"
	// Import your packages here
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// Register your services here
	fmt.Println("gRPC server running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
