package main

import (
	"context"
	"io"
	"log"

	pb "github.com/aldoclg/grpc-go/calculator/proto"
)

func doCalcMax(c pb.CalculatorServiceClient) {
	log.Println("doCalc was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while calling doCalcMax %v\n", err)
	}

	waitc := make(chan struct{})

	numbers := [10]int32{1, 5, 3, 6, 2, 20, 14, 45, 32, 47}

	go func() {
		for _, n := range numbers {
			req := &pb.NumberRequest{Number: n}

			log.Printf("Sending %d\n", n)

			stream.Send(req)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while reading the stream %v\n", err)
				break
			}

			log.Printf("doCalcMax %d\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
