package main

import (
	"context"
	"log"

	pb "github.com/aldoclg/grpc-go/calculator/proto"
)

func (s *Server) Calculate(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculate function envoked with %v\n", in)
	return &pb.CalculatorResponse{
		Result: in.A + in.B,
	}, nil
}
