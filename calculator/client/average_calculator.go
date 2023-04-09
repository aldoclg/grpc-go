package main

import (
	"context"
	"log"

	pb "github.com/aldoclg/grpc-go/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) error {
	log.Println("doAverage was invoked")
	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling doAverage %v\n", err)
	}

	var i int32
	for i = 2; i < 1000; i += 7 {
		stream.Send(&pb.NumberRequest{
			Number: i,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while calling doAverage %v\n", err)
	}

	log.Println("Average:", res.Result)
	return nil

}
