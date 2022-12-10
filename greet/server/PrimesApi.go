package main

import (
	pb "go_code/gRPCudemy/greet/proto"
	"log"
)

//
func (s *Server) PrimesApi(in *pb.PrimesRequest, stream pb.GreetService_PrimesApiServer) error {
	log.Printf("PrimesApi function was invoked with %v\n", in)
	var k int64 = 2
	for in.PrimesNum > 1 {
		if in.PrimesNum%k == 0 {
			stream.Send(&pb.PrimesResponse{
				Result: k,
			})
			in.PrimesNum = in.PrimesNum / k
		} else {
			k = k + 1
		}
	}
	return nil
}
