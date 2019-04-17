package main

import (
	"fmt"
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"../api"
)

type Server struct{}

func main() {
	lis, err := net.Listen("tcp", ":7777")

	if err != nil {
		log.Fatal("Error")
	}
	s := Server{}

	grpcServer := grpc.NewServer()

	api.RegisterPingServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error")
	}
}

func (s *Server) SayHello(ctx context.Context, in *api.PingMessage) (*api.PingMessage, error) {
	fmt.Printf("Received message: %s", in.Greeting)
	return &api.PingMessage{Greeting: "Hi there"}, nil
}
