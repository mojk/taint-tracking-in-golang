package main

import (
	"context"
	"log"
	"fmt"
	"time"
	"google.golang.org/grpc"
	pb "taint-tracking/taint-tracking"

)
const (
	address = "localhost:50051"
)

func main() {
	fmt.Println("Starting up the contorl_client..")

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect, v%", err)
	}
	defer conn.Close()
	c := pb.NewDriveClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	/* recieves the current velocity */
	rpc_getV, err := c.GetVelocity(ctx, &pb.VelocityRequest{Req: "Simple request"})
	if err != nil {
		log.Fatalf("Could not increase the velocity: %v", err)
	}
	log.Printf("Current velocity is ", rpc_getV.Velocity)

	/* increases the velocity */
	rpc_incV, err := c.IncVelocity(ctx, &pb.IncVelocityRequest{Inc: 10})
	if err != nil {
		log.Fatalf("Could not increase the velocity: %v", err)
	}
	if rpc_incV.ReturnCode == false {
		log.Printf("Could not increase the velocity")
	} else {
		log.Printf("Increasing the Velocity, response from car %v", rpc_incV.NewVelocity)
	}

	
}

