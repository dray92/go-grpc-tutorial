package main

import (
	"flag"

	"net"

	pb "github.com/dray92/go-grpc-tutorial/pb"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Implements of HelloServiceServer

type helloServer struct{}

func newHelloServer() pb.HelloServiceServer {
	return new(helloServer)
}

func (s *helloServer) Hello(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	glog.Info(msg)
	return msg, nil
}

func Run() error {
	listen, err := net.Listen("tcp", ":50051")
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
