package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/elvis-yan/grpc-demo/helloworld/helloworldpb"
	"google.golang.org/grpc"
)

const (
	addr        = ":50051"
	defaultname = "World"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	name := defaultname
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %v", r.GetMessage())

}
