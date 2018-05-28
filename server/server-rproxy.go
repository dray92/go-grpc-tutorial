package main

import (
	"flag"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/dray92/go-grpc-tutorial/internal"
	pb "github.com/dray92/go-grpc-tutorial/pb"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	helloEndpoint = flag.String(
		"hello_endpoint",
		internal.DefaultServerConfig.HostPort,
		"hello_endpoint for HelloService")
)

func RunEndPoint(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, *helloEndpoint, dialOpts)
	if err != nil {
		return err
	}

	http.ListenAndServe(address, mux)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := RunEndPoint(internal.DefaultReverseProxyConfig.Address); err != nil {
		glog.Fatal(err)
	}
}
