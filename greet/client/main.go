package main

import (
	"fmt"
	"log"

	pb "go_code/gRPCudemy/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	//Dial creates a client connection to the given target.
	//Because grpc is using SSL by defult,in case to avoid error we have to use following functions
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	//close the connection automatic when we don't need it
	defer conn.Close()

	//chose the function that you want to excute
	for {
		//registe the service client
		c := pb.NewGreetServiceClient(conn)
		var functionSelect int64
		fmt.Println("choose function.")
		fmt.Println("1:doGreet 2:doSum 3:doGreetManyTimes 4:doPrimesApi 5:doLongGreet 6:doAvgApi 7:dogreetEveryone 8:endfunction")
		fmt.Scanln(&functionSelect)
		if functionSelect == 1 {
			doGreet(c)
		} else if functionSelect == 2 {
			doSum(c)
		} else if functionSelect == 3 {
			doGreetManyTimes(c)
		} else if functionSelect == 4 {
			doPrimesApi(c)
		} else if functionSelect == 5 {
			doLongGreet(c)
		} else if functionSelect == 6 {
			doAvgApi(c)
		} else if functionSelect == 7 {
			doGreetEveryone(c)
		} else if functionSelect == 8 {
			return
		}

	}

}
