package main

import (
	"log"
	"net"

	pb "github.com/mentalclear-learningo/grpc-upnrunning/ecommerce"
	srv "github.com/mentalclear-learningo/grpc-upnrunning/service"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &srv.Server{})

	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
