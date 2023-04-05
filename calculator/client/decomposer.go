package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/aldoclg/grpc-go/calculator/proto"
)

func doDecomposition(c pb.CalculatorServiceClient) {
	log.Println("doDecomposition was invoked")

	req := &pb.DecompositionRequest{
		A: 1234567890,
	}

	stream, err := c.Decompose(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while streaming results %v\n", err)
	}

	var results string

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}

		results = fmt.Sprintf("%s %d", results, msg.Result)
	}

	log.Printf("Decompose %d in%s\n", req.A, results)

}
