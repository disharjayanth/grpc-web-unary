package main

import (
	"context"
	"fmt"
	"grpc-web/echoproto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	echoproto.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, req *echoproto.EchoRequest) (*echoproto.EchoResponse, error) {
	fmt.Println(req)
	reqMessage := req.GetMessage()
	return &echoproto.EchoResponse{
		Message: `Response from server ` + reqMessage,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalln("Error creating tcp listener:", err)
		return
	}

	s := grpc.NewServer()

	echoproto.RegisterEchoServiceServer(s, &server{})

	log.Println("Server listening @PORT:9090")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
