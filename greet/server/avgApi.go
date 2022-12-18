package main

import (
	pb "go_code/gRPCudemy/greet/proto"
	"io"
	"log"
)

func (s *Server) AvgApi(stream pb.GreetService_AvgApiServer) error {
	log.Println("AvgApi was invoked")
	var res float32
	counter := 0
	var sum float32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Current sum:%v,total numbers:%v\n", sum, counter)
			res = sum / float32(counter)
			return stream.SendAndClose(&pb.AvgResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		counter += 1
		sum += req.AvgRequestNum
		log.Printf("Reveiving: %v,index:%v\n", req, counter)
		log.Printf("Current sum:%v\n", sum)

	}
}
