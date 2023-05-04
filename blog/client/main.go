package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/aldoclg/grpc-go/blog/proto"
)

var addr string = "localhost:50053"

func main() {

	tls := true

	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Failed while loading certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	doCreateBlog(c)

}
