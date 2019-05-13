package main

/* Controls for the car
 * For every rpc-call it makes, it should send it also to the log server
 * TODO Send requests to the logserver
 * TODO
 */

import (
	"context"
	"log"
	"fmt"
	"time"
	"google.golang.org/grpc"
	pb "taint-tracking-in-golang/taint-tracking"

)
const (
	address = "localhost:50051"
)

func main() {
	fmt.Println("Starting up the control_client..")

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	conn2, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect, v%", err)
	}
	defer conn.Close()
	defer conn2.Close()

	/* connection for the driveclient */
	c := pb.NewDriveClient(conn)
	/* connection for the logclient */
	c2 := pb.NewLogClient(conn2)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	/* recieves the current velocity */
	rpc_getV, err := c.GetVelocity(ctx, &pb.VelocityRequest{Req: "Simple request"})
	if err != nil {
		log.Fatalf("Could not increase the velocity: %v", err)
	}
	log.Printf("Current velocity is ", rpc_getV.Velocity)

	//TODO
	rpc_log, err := c2.LogAction(ctx, &pb.LogRequest{Info: "Issued GetVelocity()"})
	if err != nil {
		log.Fatalf("Could not send info to the logging server: %v",err)
	}
	log.Printf("test", rpc_log.Code)

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

