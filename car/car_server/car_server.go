package main
import (
	"context"
	"log"
	"net"
	"fmt"
	"strconv"

	"google.golang.org/grpc" // grpc package
	pb "taint-tracking-in-golang/taint-tracking" // protocol buffers
)

const (
	port = ":50051" // assigned port
)

var velocity = int32(10)

type server struct{}

// Requests the current velocity of the car.
// The server will answer an VelocityRequest issued by a client-node.
// It will asnwer with an VelocityReply containing the current velocity.

func (s *server) GetVelocity(ctx context.Context, in *pb.VelocityRequest) (*pb.VelocityReply, error) {
	return &pb.VelocityReply{Velocity: velocity, Log: true}, nil
}

// Requests to increase the velocity of the car.
// The server will answer a IncVelocityRequest issued by a client-node.
// It will answer with an IncVelocityReply with the new calculated velocity and a boolean indicating if it was sucessful or not.

func (s *server) IncVelocity(ctx context.Context, in *pb.IncVelocityRequest) (*pb.IncVelocityReply, error) {
	if(velocity > 100) {
		fmt.Println("Too fast, the limit is 100 mph")
		return &pb.IncVelocityReply{ReturnCode: false}, nil
	}
	fmt.Println("Increasing the velocity with.." + strconv.Itoa(int(in.Inc)))
	velocity = IncVelocity(velocity,in.Inc)
	fmt.Println("New velocity of the car is.." + strconv.Itoa(int(velocity)))
	return &pb.IncVelocityReply{NewVelocity: velocity, ReturnCode: true}, nil
}
// Request to decrease the velocity of the car.
// The server will answer an DecVelocityRequest issued by a client-node
// It will answer with an DecVelocityReply containing the new calculated velocity and a boolean indicating if it was successful or not.
/*** FUNC ***/
/* Decreases the velocity with the value that is presented in the Request
* Answers the client with the new velocity */
func (s *server) DecVelocity(ctx context.Context, in *pb.DecVelocityRequest) (*pb.DecVelocityReply, error) {
	fmt.Println("Decreasing the velocity with.." + strconv.Itoa(int(in.Dec)))
	velocity = DecVelocity(velocity, in.Dec)
	fmt.Println("New velocity of the car is.." + strconv.Itoa(int(velocity)))
	return &pb.DecVelocityReply{NewVelocity: velocity, ReturnCode: true}, nil
}

// Function used to do the simple calculation of increasing the velocity
// @param current_velocity, int32 - the current velocity of the car
// @param increase, int32 - the value in which the car should decrease its velocity

func IncVelocity(current_velocity int32, increase int32) int32 {
	return current_velocity + increase
}

// Function used to do the simple calcuation of decreasing the velocity
// @param current_velocity, int32 - the current velocity of the car
// @param increase, int32 - the value in which the car should increase its velocity

func DecVelocity(current_velocity int32, increase int32) int32 {
	return current_velocity + increase
}

// Main function for initiating th eserver and listen to the specific port
// This server is a RegisterDriverServer, check protobuf-files for more information.

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
