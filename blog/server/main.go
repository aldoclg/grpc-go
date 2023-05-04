package main

import (
	"context"
	"log"
	"net"

	pb "github.com/aldoclg/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var collection *mongo.Collection

var addr string = "0.0.0.0:50053"

type Server struct {
	pb.BlogServiceServer
}

func main() {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://<user>:<password>@sandbox.oj51n6z.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	client, err := mongo.NewClient(opts)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on address %s", addr)

	optsTTL := []grpc.ServerOption{}
	tls := true

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates %v\n", err)
		}

		optsTTL = append(optsTTL, grpc.Creds(creds))
	}

	s := grpc.NewServer(optsTTL...)
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
