package main

import (
	"log"
	"time"

	gRPC_testing "github.com/jskoven/gRPC_testing/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("192.168.43.169:9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := gRPC_testing.NewChatServiceClient(conn)

	message := gRPC_testing.Message{
		Body:    "Det her er min besked til dig Maltus",
		TimeNow: time.Now().String(),
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling sayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Body)
	log.Printf("Time on server is: %s", &message.TimeNow)
}
