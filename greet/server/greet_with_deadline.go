package main

import (
	"context"
	"log"
	"time"

	pb "github.com/aldoclg/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function envoked with %v\n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client cancelled the request")
			return nil, status.Error(codes.Canceled, "Request canceled")
		}
	}

	time.Sleep(1 * time.Second)

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
