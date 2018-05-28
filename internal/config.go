package internal

import "fmt"

import "github.com/dray92/go-grpc-tutorial/internal/server"

type ServerConfig struct {
	HostPort string
}

type ReverseProxyConfig struct {
	Address string
}

// not sure if this is the ideal place for this....
var DefaultServerConfig = ServerConfig{
	HostPort: fmt.Sprintf("localhost%s", server.DefaultGRPCServerConfig.Address),
}

var DefaultReverseProxyConfig = ReverseProxyConfig{
	Address: ":8080",
}
