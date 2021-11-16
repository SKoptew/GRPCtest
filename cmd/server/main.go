package main

import (
	"google.golang.org/grpc"
	"grpctest/pkg/adder"
	"grpctest/pkg/api"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("net.Listen failed:", err)
	}

	s := grpc.NewServer()
	api.RegisterAdderServer(s, &adder.AdderServer{})

	log.Print("server listening at", listener.Addr())
	if err = s.Serve(listener); err != nil {
		log.Fatal("s.Serve failed:", err)
	}
}