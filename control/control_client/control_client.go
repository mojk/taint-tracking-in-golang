package main

/* Controls for the car
 * For every rpc-call it makes, it should send it also to the log server
 * TODO Send requests to the logserver - DONE (separate func)
 */

import (
	"context"
	"log"
	"fmt"
	"time"
//	"reflect"

	"google.golang.org/grpc"
	pb "taint-tracking-in-golang/taint-tracking"

)
/* addresses and the ports used for communicating with the services */
const (
	address_car = "localhost:50051"
	address_log = "localhost:50052"
)

// this func will log the event performed in the form of a remote procedure call
func log_event(information string, client pb.LogClient, ctx context.Context) {

	rpc_log, err := client.LogAction(ctx, &pb.LogRequest{Info: information})
	if err != nil {
		log.Fatalf("Could not send info to the logging server: %v",err)
	}
	log.Printf("Sucessful? %v", rpc_log.Code)
}
// this func will be send to the control_server and ask if it should filter any data before sending it to the log_server or something
//TODO
func filter_event() {
	rpc_filter, err := client.FilterData(ctx, &pb.FilterRequest{  })

}

func main() {
	fmt.Println("Starting up the control_client..")

	/* Setting up two connections for two different ports */
	conn_car, err := grpc.Dial(address_car, grpc.WithInsecure())
	conn_log, err := grpc.Dial(address_log, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect,%v", err)
	}
	defer conn_car.Close()
	defer conn_log.Close()

	/* Creating the clients for respective services */
	c := pb.NewDriveClient(conn_car)
	c2 := pb.NewLogClient(conn_log)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()	

	/*** RPC CALL ***/
	rpc_getV, err := c.GetVelocity(ctx, &pb.VelocityRequest{Req: "Simple request"})
	if err != nil {
		log.Fatalf("Could not increase the velocity: %v", err)
	}
	if (rpc_getV.Log == true) { //If the request was succesful I think?
		log_event("GetVelocity()", c2, ctx)
		log.Printf("Current velocity is %v", rpc_getV.Velocity)
	}

	/*** RPC CALL ***/
	rpc_incV, err := c.IncVelocity(ctx, &pb.IncVelocityRequest{Inc: 10})
	if err != nil {
		log.Fatalf("Could not increase the velocity: %v", err)
	}
	if rpc_incV.ReturnCode == false {
		log.Printf("Could not increase the velocity")
	} else {
		log_event("IncVelocity()", c2, ctx)
		log.Printf("Increasing the velocity, current velocity = %v", rpc_incV.NewVelocity)
	}	
}
