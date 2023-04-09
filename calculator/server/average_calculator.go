package main

import (
	"io"
	"log"

	pb "github.com/aldoclg/grpc-go/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average function envoked")

	sum := 0
	i := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Calc: %d / %d", sum, i)
			return stream.SendAndClose(&pb.AverageResponse{
				Result: float32(sum / i),
			})
		}

		if err != nil {
			log.Printf("Error while reading the stream %v\n", err)
			return err
		}

		log.Printf("[Average] Number: %d\n", req.Number)

		sum += int(req.Number)
		i++
	}

}
