package main

import (
	"../api"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct{}

func main() {
	lis, err := net.Listen("tcp", ":7777")

	if err != nil {
		log.Fatal("Error")
	}
	s := Server{}

	grpcServer := grpc.NewServer()

	api.RegisterAddServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error")
	}
}

func (s *Server) AddNumbers(ctx context.Context, in *api.Request) (*api.Response, error) {
	// log.Printf("Received message: %s", in.Greeting)
	res := calculate(in.A, in.B)

	log.Printf("Calculate: %d", res)
	return &api.Response{Result: res}, nil
}

func calculate(a, b int32) int32 {
	return a + b
}
