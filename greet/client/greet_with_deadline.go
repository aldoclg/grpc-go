package main

import (
	"context"
	"log"
	"time"

	pb "github.com/aldoclg/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{
		FirstName: "D",
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Deadline exceeded %v\n", e)
			} else {
				log.Fatalf("Could not greet: %v\n", err)
			}

		} else {
			log.Fatalf("Could not greet: %v\n", err)
		}
	}

	log.Printf("Greeting: %s\n", res.Result)
}
