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

	for {
		var first, second int32

		fmt.Printf("First: ")
		fmt.Scanf("%d", &first)

		fmt.Printf("Second: ")
		fmt.Scanf("%d", &second)

		//response, err := c.SayHello(context.Background(), &api.Request{A: first, B: second})
		c.AddNumbers(context.Background(), &api.Request{A: first, B: second})

		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}

		//fmt.Printf("The sum of %d and %d are %d\n", first, second, response.Result)
	}

}
