package main

import (
	"fmt"
	"log"

	pb "github.com/aldoclg/grpc-go/greet/proto"
)

func (s *Server) GreetManyTimes(req *pb.GreetRequest, service pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes was invoked %v\n", req)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", req.FirstName, i)

		service.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
