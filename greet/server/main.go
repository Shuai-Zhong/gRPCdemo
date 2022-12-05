package main

import (
	pb "go_code/gRPCudemy/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
	//pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Fail to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer()
	//register the GreetServiceServer
	//the grpc server need an instance for the Greet service
	//we use "Server" type to implement all the RPC endpoints
	pb.RegisterGreetServiceServer(s, &Server{})
	//pb.RegisterSumServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve:%v\n", err)
	}
}
