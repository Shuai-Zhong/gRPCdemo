package main

import (
	pb "go_code/gRPCudemy/greet/proto"
	"log"
)

func (s *Server) PrimesApi(in *pb.PrimesRequest, stream pb.GreetService_PrimesApiServer) error {
	log.Printf("PrimesApi function was invoked with %v\n", in)
	return nil
}
