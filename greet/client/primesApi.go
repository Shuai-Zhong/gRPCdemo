package main

import (
	"context"
	"fmt"
	pb "go_code/gRPCudemy/greet/proto"
	"io"
	"log"
)

func doPrimesApi(c pb.GreetServiceClient) {
	log.Println("doPrimesApi was invoked")
	fmt.Println("pls enter a number")
	var ReqNum int64
	fmt.Scanln(&ReqNum)
	//passing value to server
	req := &pb.PrimesRequest{
		PrimesNum: ReqNum,
	}
	stream, err := c.PrimesApi(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling doPrimesApi:%v\n", err)
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
		log.Printf("doPrimesApi res:%v\n", msg.Result)
	}

}
