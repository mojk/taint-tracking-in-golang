package main

import (
	"fmt"
	"context"
	"log"
	"os"
	"time"
	"google.golang.org/grpc"
	pb "taint-tracking/taint-tracking"

)
const (
	address = "localhost:50051"
	defaultName = "car-client"
)

func main() {
	fmt.Println("Starting the car-client.. VRUUM VRUUUM")

	/* Setting up a connection to the Server  */
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v",err)
	}
	
}
