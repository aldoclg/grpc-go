package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/aldoclg/grpc-go/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {

	log.Println("GreetEveryone was invoked")

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

		err = stream.Send(&pb.GreetResponse{
			Result: fmt.Sprintf("Hello %s", req.FirstName),
		})

		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}

		log.Printf("[GreetEveryone] First name: %s\n", req.FirstName)

	}
}
