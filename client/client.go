package main

import (
	"log"
	"os"

	"fmt"

	pb "github.com/dray92/go-grpc-tutorial/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "debo"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	msg := fmt.Sprintf("Hello %v!", name)

	r, err := c.Hello(context.Background(), &pb.Message{Id: "1", Msg: msg})
	if err != nil {
		log.Fatalf("Oops! Something went wrong: %v", err)
	}
	log.Printf(r.Msg)
}
