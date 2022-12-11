package main

import (
	"context"
	pb "go_code/gRPCudemy/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "Marie"},
		{FirstName: "Clement"},
		{FirstName: "Test"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet %v\n", err)
	}

	//sending requests
	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		//1 sec delay after sending a request
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}
	log.Printf("LongGreet result: %s\n", res.Result)

}
