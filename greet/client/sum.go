package main

import (
	"context"
	"fmt"
	pb "go_code/gRPCudemy/greet/proto"
	"log"
)

func doSum(c pb.GreetServiceClient) {
	log.Println("doSum was invoked")
	//send a request with the numbers we want to use for function sum()
	var UserInputFirst int64
	var UserInputSec int64
	fmt.Println("pls input first num")
	fmt.Scanln(&UserInputFirst)
	fmt.Println("pls input sec num")
	fmt.Scanln(&UserInputSec)
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNum: UserInputFirst,
		SecNum:   UserInputSec,
	})

	if err != nil {
		log.Fatalf("could not sum: %v\n", err)
	}
	log.Println("Sum:", res.Result)
}
