package main

import (
	"context"
	"log"

	pb "github.com/aldoclg/grpc-go/calculator/proto"
)

func doCalc(c pb.CalculatorServiceClient) {
	log.Println("doCalc was invoked")

	res, err := c.Calculate(context.Background(), &pb.CalculatorRequest{
		A: 1,
		B: 4,
	})

	if err != nil {
		log.Fatalf("Could not calc: %v\n", err)
	}

	log.Printf("Result: %d\n", res.Result)
}
