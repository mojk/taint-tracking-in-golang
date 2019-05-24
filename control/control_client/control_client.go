package main

/* Controls for the car
 * For every rpc-call it makes, it should send it also to the log server
 */

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	pb "taint-tracking-in-golang/taint-tracking"
)

/* addresses and the ports used for communicating with the services */
const (
	address_car     = "localhost:50051"
	address_log     = "localhost:50052"
	address_control = "localhost:50053"
)

// var
var filter_get bool
var filter_inc bool
var filter_dec bool
var current_velocity int32

//
func get_velocity(client pb.DriveClient, ctx context.Context, logs pb.LogClient) {
	rpc_getV, err := client.GetVelocity(ctx, &pb.VelocityRequest{Req: "Simple request"})
	if err != nil {
		log.Fatalf("Could not increase the velocity: %v", err)
	}
	if rpc_getV.Log == true {
		log_event("GetVelocity()", logs, ctx) // if successful, send info to the logging server
		current_velocity = rpc_getV.Velocity
		log.Printf("Current velocity is %v", current_velocity)
	}
}

// function for decreasing the velocity of the car
func decrease_velocity(speed int, client pb.DriveClient, ctx context.Context, logs pb.LogClient) {
	rpc_decV, err := client.DecVelocity(ctx, &pb.DecVelocityRequest{Dec: 40})
	if err != nil {
		log.Fatalf("Could not decrease the velocity: %v", err)

	}
	if rpc_decV.ReturnCode == false {
		log.Printf("Could not decrease the velocity")
	} else {
		log_event("DecVelocity()", logs, ctx) // if successful, send info to the logging server
		current_velocity = rpc_decV.NewVelocity
		log.Printf("Decreasing the velocity, current velocity = %v", current_velocity)
	}
}

// function for increasing the velocity of the car
func increase_velocity(speed int, client pb.DriveClient, ctx context.Context, logs pb.LogClient) {

	rpc_incV, err := client.IncVelocity(ctx, &pb.IncVelocityRequest{Inc: 10}) // sending a request to increase the velocity
	if err != nil {
		log.Fatalf("Could not increase the velocity: %v", err)
	}
	if rpc_incV.ReturnCode == false {
		log.Printf("Could not increase the velocity")
	} else {
		log_event("IncVelocity()", logs, ctx) // if successful, send info to the logging server
		current_velocity = rpc_incV.NewVelocity
		log.Printf("Increasing the velocity, current velocity = %v", current_velocity)
	}
}

/*
This function will send a string to the log_server who in turn will display every rpc-call that the contol has made
*/
func log_event(information string, client pb.LogClient, ctx context.Context) {
	fmt.Println("Sending info to the logging server")
	rpc_log, err := client.LogAction(ctx, &pb.LogRequest{Info: information})
	if err != nil {
		log.Fatalf("Could not send info to the logging server: %v", err)
	}

	if rpc_log.Code {
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
	if rpc_filter.Action == true {
		fmt.Println("Changing the filtering-options..")

		if rpc_filter.Inc == true {
			filter_inc = true
		}

		if rpc_filter.Dec == true {
			filter_dec = true
		}

		if rpc_filter.Get == true {
			filter_get = true
		}
	}
}

func remote_source() int {
	return 10
}

func main() {
	fmt.Println("Starting up the control_client..")
	current_velocity = 0

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

	for {

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		/* Checking filter options before performing rpc */
		filter_event(c3, ctx)

		// Requesting the velocity to decide what it should do
		get_velocity(c, ctx, c2)

		// Decision made depending on the current velocity
		if current_velocity < 100 { // if the velocity is below 100 we can increase!

			// Increase velocity with this new neat function that uses machine learning lolz
			increase_velocity(remote_source(), c, ctx, c2)
		} else if current_velocity >= 100 { // if the velocity is higher we should decrease
			decrease_velocity(10, c, ctx, c2)
		}
	}
}
