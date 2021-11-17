package main

import (
	"context"
	"google.golang.org/grpc"
	"grpctest/pkg/api"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("cannot connect:", err)
	}
	defer conn.Close()

	client := api.NewAdderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	response, err := client.Add(ctx, &api.AddRequest{X: 1, Y: 4})
	if err != nil {
		log.Fatal("cannot remote call Add", err)
	}
	log.Print("result:", response.GetResult())
}
