package main

import (
  "log"
  "../api"
  "context"
  "google.golang.org/grpc"
  "bufio"
  "fmt"
  "os"
)

func main() {

    reader := bufio.NewReader(os.Stdin)
    var conn *grpc.ClientConn

    conn, err := grpc.Dial(":7777", grpc.WithInsecure())

    if err != nil {
        log.Fatalf("did not connect: %s", err)
    }

    defer conn.Close()

    c := api.NewPingClient(conn)

    for {
        fmt.Printf("Enter your message: ")

        data, _ := reader.ReadString('\n')

        response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: data})

        if err != nil {
            log.Fatalf("Error when calling SayHello: %s", err)
        }

        fmt.Printf("Response from server: %s\n", response.Greeting)
    }

}
