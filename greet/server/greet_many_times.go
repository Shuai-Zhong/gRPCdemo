package main

import (
	"fmt"
	pb "go_code/gRPCudemy/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with:%v\n", in)
	//return 10 times "Hello Clement, counter"
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s,number %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
