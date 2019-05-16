package main

/*
 * The Control_server will communicate with the control_client & the log_client
 * It will listen on por 50053 for incoming requests
 * control_client will send requests if it should filter out data before sending it
 * to the log_server
 * 
 * log_client will send requests to the control_server to filter out different data
 */
 
import (
	"context"
	"log"
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "taint-tracking-in-golang/taint-tracking"

)

const (
	port = ":50053"
)

var getVel bool
var incVel bool
var decVel bool

type server struct {}

/* control_client issues these request*/
/* asks for what it should filter or not */
func (s *server) FilterQuestion(ctx context.Context, in *pb.FilterQuestionRequest) (*pb.FilterQuestionReply, error) {
	return &pb.FilterQuestionReply{Action: true, Get: getVel, Inc: incVel, Dec: decVel}, nil
}

/* log_client issues these requests*/
/* */
func (s *server) FilterData(ctx context.Context, in *pb.FilterRequest) (*pb.FilterReply, error) {
	//inital values
	getVel = true
	incVel = true
	decVel = true

	if (in.GetVel == true) {
		fmt.Println("Filtering out GetVelocity()")
		getVel = false
	} else if (in.IncVel == true) {
		fmt.Println("Filtering out IncVelocity()")
		incVel = false
	} else if (in.DecVel == true) {
		fmt.Println("Filtering out DecVelocity()")
		decVel = false
	}
	
	return &pb.FilterReply{Success: true, GetVel: getVel, IncVel: incVel, DecVel: decVel},nil
}

func main() {
	fmt.Println("Starting up the control_server")
	
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen", port)
	}
	s := grpc.NewServer()
	pb.RegisterFilterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server", err)
	}
}
