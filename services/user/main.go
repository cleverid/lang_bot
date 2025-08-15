package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"user/app"
	"user/contracts"
)

func main() {
	lis, err := net.Listen("tcp", ":15000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	contracts.RegisterUserServer(s, app.New())

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
