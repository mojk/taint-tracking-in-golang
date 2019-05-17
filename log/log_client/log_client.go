package main

import (
	"context"
	"log"
	"fmt"
	"time"
	"strconv"

	"google.golang.org/grpc"
	pb "taint-tracking-in-golang/taint-tracking"

)

const (
	address_control = "localhost:50053"
)

func main() {
	fmt.Println("Launching the log_client...")

	/* setting up the connection  */
	conn_control, err := grpc.Dial(address_control, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect, %v", err)
	}
	defer conn_control.Close() 
	
	/* creating the client */
	c := pb.NewFilterClient(conn_control)
	
	for {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//rpc-call
	rpc_filter, err := c.FilterData(ctx, &pb.FilterRequest{GetVel: false, IncVel: true, DecVel: false})
	if err != nil {
		log.Fatalf("Could not request filtering of data %v", err)
	}
	log.Printf("Filter request has been sent! %v" +  strconv.FormatBool(rpc_filter.Success))

	}
}
