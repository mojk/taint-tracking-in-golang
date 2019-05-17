package main

import (
	"../api"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	c := api.NewAddClient(conn)

	a := fun_grpc(c)

	writer(a)

}

func reader() int32 {
	return 10
}

func writer(s int32) {

}

func fun_grpc(c api.AddClient) int32 {
	response, err := c.AddNumbers(context.Background(), &api.Request{A: 1, B: 1})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	fmt.Printf("The sum of %d and %d are %d\n", 1, 1, response.Result)

	return response.Result
}
