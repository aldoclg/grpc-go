package main

import (
	"context"
	"log"

	"math/rand"
	"time"

	pb "github.com/aldoclg/grpc-go/blog/proto"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz"
)

func doCreateBlog(c pb.BlogServiceClient) string {
	log.Println("The method doCreateBlog was invoked")

	rand.Seed(time.Now().UnixNano())

	// Getting random character
	letter := charset[rand.Intn(len(charset))]

	letterString := string(letter)

	blog := &pb.Blog{
		AuthorId: letterString,
		Title:    letterString,
		Content:  letterString,
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("An error has happened %s\n", err.Error())
	}

	log.Printf("Blog has been created")

	return res.Id
}
