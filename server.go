package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %t", err)
	}
	grpcSever := grpc.NewServer()
	if err := grpcSever.Serve(lis); err != nil {
		log.Fatalf("Failed")
	}

}
