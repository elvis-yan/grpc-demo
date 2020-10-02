package main

import (
	"context"
	"log"
	"net"

	pb "github.com/elvis-yan/grpc-demo/helloworld/helloworldpb"
	"google.golang.org/grpc"
)

const (
	addr = ":50051"
)

type helloWorldServer struct{}

func (s *helloWorldServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &helloWorldServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
