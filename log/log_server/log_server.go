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
	port = ":50052"
)

type server struct{}

func (s *server) GetVelocity() {

}
func (s *server) IncVelocity() {

}
func (s *server) DecVelocity() {

}

func main() {
	fmt.Println("Starting the log_server..")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen to port: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDriveServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
