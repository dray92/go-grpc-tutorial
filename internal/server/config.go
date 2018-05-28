package server

type GRPCServerConfig struct {
	Network, Address string
}

// grpc server will serve at 50000
var DefaultGRPCServerConfig = GRPCServerConfig{
	Network: "tcp",
	Address: ":50000",
}
