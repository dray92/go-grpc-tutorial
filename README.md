# Golang gRPC Tutorial + Swagger JSON
Basic tutorial into setting up a gRPC reverse proxy that will
transform REST JSON into gRPC, along with instructions on how
to generate Swagger definitions.

## Goals
- service defintion in `service.proto`
- generate client and server code using [`protoc`](https://github.com/google/protobuf).
- wrap generated code to expose light client (and server).
- generate gateway code for HTTP/JSON reverse proxy.

## TODO
- get Swagger UI to render service's Swagger JSON.

## Basic Setup
- use [`gvm`](https://github.com/moovweb/gvm) or something to get Go >= 1.6
- `brew install protobuf` or `apt install libprotobuf-dev`
- `$ go help gopath # make sure GOPATH is set`

### Golang installations
```bash
go get google.golang.org/grpc
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

## Architecture
This is the basic architecture of a grpc-gateway application.

![Architecture](https://camo.githubusercontent.com/e75a8b46b078a3c1df0ed9966a16c24add9ccb83/68747470733a2f2f646f63732e676f6f676c652e636f6d2f64726177696e67732f642f3132687034435071724e5046686174744c5f63496f4a707446766c41716d35774c513067677149356d6b43672f7075623f773d37343926683d333730 "Architecture")

### Service definition
TBD

