package main

/* Controls for the car
 * For every rpc-call it makes, it should send it also to the log server
 * TODO Send requests to the logserver - DONE (separate func)
 * TODO send requests to the control_server - DONE (separate func)
 */

import (
	"context"
	"log"
	"fmt"
	"time"
	"strconv"

	"google.golang.org/grpc"
	pb "taint-tracking-in-golang/taint-tracking"

)

/* addresses and the ports used for communicating with the services */

const (
	address_car = "localhost:50051"
	address_log = "localhost:50052"
	address_control = "localhost:50053"
)

var filter_get bool
var filter_inc bool
var filter_dec bool

/*
This function will send a string to the log_server who in turn will display every rpc-call that the contol has made
*/

func log_event(information string, client pb.LogClient, ctx context.Context) {
	fmt.Println("Sending info to the logging server")
	rpc_log, err := client.LogAction(ctx, &pb.LogRequest{Info: information})
	if err != nil {
		log.Fatalf("Could not send info to the logging server: %v",err)
	}
	
	if (rpc_log.Code) {
		log.Printf("Sucessful! %v:", rpc_log.Code)		
	} else {
		log.Printf("Failure! %v:", rpc_log.Code)		
	}
}

/* 
This function will perform a rpc to the control_server asking what the filter options look like. The parameters are of the type bool and true indicates that it should be filtered, and false indicates that is should be kept.
@param get - GetVelocity()
@param inc - IncVelocity()
@param dec - DecVelocity()
*/

func filter_event(client pb.FilterClient, ctx context.Context) {
	rpc_filter, err := client.FilterQuestion(ctx, &pb.FilterQuestionRequest{Action: true})
	if err != nil {
		log.Fatalf("Could not request filtering from the control_server: %v", err)
	}
	log.Printf("Response from the server %v", strconv.FormatBool(rpc_filter.Action))

	// If it returns true, it means the filter options may be changed
	if (rpc_filter.Action == true) {
		fmt.Println("Changing the filtering-options..")

		if(rpc_filter.Inc == true) {
			filter_inc = true
		}

		if(rpc_filter.Dec == true) {
			filter_dec = true
		}

		if(rpc_filter.Get == true) {
			filter_get = true
		}
	}
}

func main() {
	fmt.Println("Starting up the control_client..")

	/* Setting up two connections for two different ports */
	conn_car, err := grpc.Dial(address_car, grpc.WithInsecure())
	conn_log, err := grpc.Dial(address_log, grpc.WithInsecure())
	conn_control, err := grpc.Dial(address_control, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect,%v", err)
	}
	defer conn_car.Close()
	defer conn_log.Close()
	defer conn_control.Close()

	/* Creating the clients for respective services */
	c := pb.NewDriveClient(conn_car)
	c2 := pb.NewLogClient(conn_log)
	c3 := pb.NewFilterClient(conn_control)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()	

	/* Checking filter options before performing rpc */
	filter_event(c3, ctx);


	/*** RPC CALL - Get the velocity ***/
	rpc_getV, err := c.GetVelocity(ctx, &pb.VelocityRequest{Req: "Simple request"})
	if err != nil {
		log.Fatalf("Could not increase the velocity: %v", err)
	}
	if (rpc_getV.Log == true) {
		log_event("GetVelocity()", c2, ctx) // if successful, send info to the logging server
		log.Printf("Current velocity is %v", rpc_getV.Velocity)
	}

	/*** RPC CALL - Increasing the velocity ***/
	rpc_incV, err := c.IncVelocity(ctx, &pb.IncVelocityRequest{Inc: 10})
	if err != nil {
		log.Fatalf("Could not increase the velocity: %v", err)
	}
	if rpc_incV.ReturnCode == false {
		log.Printf("Could not increase the velocity")
	} else {
		log_event("IncVelocity()", c2, ctx) // if successful, send info to the logging server
		log.Printf("Increasing the velocity, current velocity = %v", rpc_incV.NewVelocity)
	}	
}
