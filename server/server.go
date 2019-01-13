package main

import (
	"flag"
	"fmt"

	"net"

	"github.com/dray92/go-grpc-tutorial/internal/server"
	pb "github.com/dray92/go-grpc-tutorial/pb"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Implementation of HelloServiceServer

type helloServer struct{}

func newHelloServer() pb.HelloServiceServer {
	return new(helloServer)
}

func (s *helloServer) Hello(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	glog.Info(msg)
	fmt.Printf("Received message: %+v\n", msg)

	if alternativeAuth := msg.GetVerySecureAlternativeAuth(); alternativeAuth != "" {
		fmt.Println("picked up alternative auth from oneof")
	} else if auth := msg.GetAuth(); auth != nil {
		fmt.Println("picked up traditional auth from oneof")
		auth.
	}

	ack := pb.Message{
		Id: msg.Id,
		Msg: fmt.Sprintf("Acknowledging message: \"%s\"!", msg.Msg),

	}
	return &ack, nil
}

func Run() error {
	listen, err := net.Listen(server.DefaultGRPCServerConfig.Network, server.DefaultGRPCServerConfig.Address)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterHelloServiceServer(server, newHelloServer())
	server.Serve(listen)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := Run(); err != nil {
		glog.Fatal(err)
	}
}
