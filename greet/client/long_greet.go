package main

import (
	"context"
	"log"

	pb "github.com/aldoclg/grpc-go/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Test client stream 1"},
		{FirstName: "Test client stream 2"},
		{FirstName: "Test client stream 3"},
		{FirstName: "Test client stream 4"},
	}

	stream, err := c.LongGreet(context.Background())

	for _, req := range reqs {
		log.Printf("Req: %v\n", req)
		err := stream.Send(req)

		if err != nil {
			log.Fatalf("Error while calling doLongGreet %v - %v\n", req, err)
		}
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while calling doLongGreet %v\n", err)
	}

	log.Printf("Response %s\n", res.Result)
}
