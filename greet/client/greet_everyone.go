package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/aldoclg/grpc-go/greet/proto"
)

func doGreetEveyone(c pb.GreetServiceClient) {
	log.Println("doGreetEveyone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling doLongGreet %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for i := 0; i < 20; i++ {
			req := &pb.GreetRequest{FirstName: fmt.Sprintf("Bidirecional %d", i)}

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

			log.Printf("doGreetEveyone %s\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
