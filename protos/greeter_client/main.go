package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "github.com/seekplum/plum/protos/helloworld"
)

const (
	address     = "localhost:50051"
	defaultName = "World"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dit not connect: %s", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("cloud not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
