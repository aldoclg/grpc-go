package main

import (
	"log"

	pb "github.com/aldoclg/grpc-go/calculator/proto"
)

func (s *Server) Decompose(req *pb.DecompositionRequest, service pb.CalculatorService_DecomposeServer) error {
	log.Printf("Decompose was invoked %v\n", req)
	number := req.A
	var div int32

	div = 2

	for number > 1 {
		if (number % div) == 0 {
			number = number / div
			res := &pb.CalculatorResponse{
				Result: div,
			}
			err := service.Send(res)
			if err != nil {
				return err
			}
		} else {
			div += 1
		}
	}
	return nil
}
