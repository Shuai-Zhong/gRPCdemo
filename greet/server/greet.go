package main

import (
	"context"
	pb "go_code/gRPCudemy/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "hello " + in.FirstName,
	}, nil
}
