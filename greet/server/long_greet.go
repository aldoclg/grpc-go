package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/aldoclg/grpc-go/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet was invoked \n")

	res := ""

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Printf("Error while reading the stream %v\n", err)
			return err
		}

		log.Printf("[LongGreet] Frst name: %s\n", msg.FirstName)

		res += fmt.Sprintf("Hello %s\n", msg.FirstName)
	}

}
