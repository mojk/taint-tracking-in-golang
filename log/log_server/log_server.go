package main

/* Service used for logging the actions of the controls 
 *
 * For every LogAction the service recieves, it will save and display the information
 * it has recieved
 */

import (
	"context"
	"log"
	"net"
	"fmt"

	"google.golang.org/grpc"
	pb "taint-tracking-in-golang/taint-tracking"
)

const (
	port = ":50052"
)

type server struct{}

/* control_client will issue this request */
func (s *server) LogAction(ctx context.Context, in *pb.LogRequest) (*pb.LogReply, error) {
	fmt.Println("New Log! " + in.Info)
	return &pb.LogReply{Code: true}, nil
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
