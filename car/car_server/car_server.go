package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"strconv"

	"google.golang.org/grpc" // grpc package
	pb "taint-tracking/taint-tracking" // protocol buffers
)

/* constants*/
const (
	port = ":50051"
)

/* variables */
var velocity = int32(10)

type server struct{}

/* returning the current velocity of the car */

func (s *server) GetVelocity(ctx context.Context, in *pb.VelocityRequest) (*pb.VelocityReply, error) {
	return &pb.VelocityReply{Velocity: velocity}, nil
}

/* Increases the velocity with the value that is presented in the Request 
 * Answers the client with the new velocity
 */

func (s *server) IncVelocity(ctx context.Context, in *pb.IncVelocityRequest) (*pb.IncVelocityReply, error) {
	if(velocity > 100) {
		fmt.Println("Too fast, the limit is 100 mph")
		return &pb.IncVelocityReply{ReturnCode: false}, nil
	}
	fmt.Println("Increasing the velocity with.." + strconv.Itoa(int(in.Inc)))
	velocity = IncVelocity(velocity,in.Inc)
	fmt.Println("New velocity of the car is.." + strconv.Itoa(int(velocity)))
	return &pb.IncVelocityReply{NewVelocity: velocity}, nil
}

/* Decreases the velocity with the value that is presented in the Request
 * Answers the client with the new velocity
 */
func (s *server) DecVelocity(ctx context.Context, in *pb.DecVelocityRequest) (*pb.DecVelocityReply, error) {
	fmt.Println("Decreasing the velocity with.." + strconv.Itoa(int(in.Dec)))
	velocity = DecVelocity(velocity, in.Dec)
	fmt.Println("New velocity of the car is.." + strconv.Itoa(int(velocity)))
	return &pb.DecVelocityReply{NewVelocity: velocity}, nil
}

/* function for doing the simple aritmethics for increasing the velocity */
func IncVelocity(current_velocity int32, increase int32) int32 {
	return current_velocity + increase
}

func DecVelocity(current_velocity int32, increase int32) int32 {
	return current_velocity + increase
}

func main() {
	fmt.Println("Starting up the car server..")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDriveServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
