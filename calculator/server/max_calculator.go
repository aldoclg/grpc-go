package main

import (
	"io"
	"log"
	"time"

	pb "github.com/aldoclg/grpc-go/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max was invoked")

	array := make([]int32, 0)
	var last int32

	for {
		req, err := stream.Recv()

		if req == nil {
			continue
		}

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Printf("Error while reading the stream %v\n", err)
			return err
		}

		log.Println("Received", req.Number)

		time.Sleep(2 * time.Second)

		if last <= req.Number {

			array = append(array, req.Number)

			err := stream.Send(&pb.CalculatorResponse{
				Result: req.Number,
			})

			if err != nil {
				log.Fatalf("Error while reading the stream %v\n", err)
			}
		}
		last = req.Number
	}
}
