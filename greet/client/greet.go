package main

import (
	"context"
	"fmt"
	pb "go_code/gRPCudemy/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	//send a request with firstname
	fmt.Println("what's your name")
	var UserInputName string
	fmt.Scanln(&UserInputName)
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: UserInputName,
	})

	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", res.Result)
}
