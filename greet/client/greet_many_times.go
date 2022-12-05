package main

import (
	"context"
	pb "go_code/gRPCudemy/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")
	//passing value to server
	req := &pb.GreetRequest{
		FirstName: "Shuai Zhong",
	}
	//create a stream for function GreetManyTimes
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes:%v\n", err)
	}

	for {
		//Recv() can allow us to accept return values from server side by multiple times
		msg, err := stream.Recv()
		//meaning the process of accepting values is end
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream :%v\n", err)
		}
		log.Printf("GreetManyTime:%s\n", msg.Result)
	}

}
