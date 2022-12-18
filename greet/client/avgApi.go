package main

import (
	"context"
	"fmt"
	pb "go_code/gRPCudemy/greet/proto"
	"log"
)

func doAvgApi(c pb.GreetServiceClient) {
	log.Println("doAvgApi was invoked")
	stream, err := c.AvgApi(context.Background())
	if err != nil {
		log.Fatalf("Error while calling AvgApi %v\n", err)
	}
	for {
		fmt.Println("choose a function 1.enter a number to AvgAPi,2.stop the process")
		UserInput := 0
		fmt.Scanln(&UserInput)
		if UserInput == 1 {
			var numberTosend float32
			fmt.Println("please enter the number that you want to send")
			fmt.Scanln(&numberTosend)
			log.Printf("sending number:%v\n", numberTosend)
			stream.Send(&pb.AvgRequest{
				AvgRequestNum: numberTosend,
			})
		} else if UserInput == 2 {
			res, err := stream.CloseAndRecv()
			if err != nil {
				log.Fatalf("Error while receiving response from AvgApi: %v\n", err)
			}
			log.Printf("AvgApi result: %v\n", res.Result)
			return
		}
	}

}
