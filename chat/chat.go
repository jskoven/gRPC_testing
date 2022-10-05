package gRPC_testing

import (
	"log"
	"time"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	log.Printf("Received time from client: %s", message.TimeNow)
	return &Message{Body: "Server received message and time. Server time is: ", TimeNow: time.Now().String()}, nil
}
