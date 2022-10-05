package main

import (
	"bufio"
	"log"
	"os"
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

	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()
	textToSend := Scanner.Text()

	message := gRPC_testing.Message{
		Body:    textToSend,
		TimeNow: time.Now().String(),
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling sayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Body)
	log.Printf("Time on server is: %s", response.TimeNow)
	timeIntime, err := time.Parse("", response.TimeNow)
	if err != nil {
		log.Printf("error in parsing time: %s", err)
	}
	log.Printf("Time parsed is: ", timeIntime)
}
