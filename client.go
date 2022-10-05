package main

import (
	"log"

	gRPC_testing "github.com/jskoven/gRPC_testing/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := gRPC_testing.NewChatServiceClient(conn)

	message := gRPC_testing.Message{
		Body: "hello from the client yee!!",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling sayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Body)
}
