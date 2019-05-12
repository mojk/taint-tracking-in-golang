package main

import (
	"context"
	"log"
	"net"
	"fmt"

	"google.golang.org/grpc"
	pb "taint-tracking/taint-tracking"
)

const (
	port = ":50051"
)

type server struct{}

/* server will handle logrequests*/

func (s *server) LogAction(ctx context.Context, in *pb.LogRequest) (*pb.LogReply, error) {
	return &pb.LogReply{Code: "Logging done"}, nil
}

func main() {
	fmt.Println("Starting the log_server..")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen to port: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLogServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
