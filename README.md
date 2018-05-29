# gRPC Tutorial + Swagger JSON
Basic tutorial into setting up a gRPC reverse proxy that will
transform REST JSON into gRPC, along with instructions on how
to generate Swagger definitions.
Instructions on how to run client/server in other languages.

## Goals
- service defintion in `service.proto`
- generate client and server code using [`protoc`](https://github.com/google/protobuf).
- wrap generated code to expose light client (and server).
- generate gateway code for HTTP/JSON reverse proxy.

## TODO
- get Swagger UI to render service's Swagger JSON.
- investigate python reverse proxy support [#398](https://github.com/grpc-ecosystem/grpc-gateway/issues/398)
- write up instructions for java {[grpc-jersey](https://github.com/fullcontact/grpc-jersey)}
- find a contributor for nodejs

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
This is the basic architecture of a grpc-gateway application {from [`grpc-ecosystem/grpc-gateway/`](https://github.com/grpc-ecosystem/grpc-gateway/)}.

![Architecture](https://camo.githubusercontent.com/e75a8b46b078a3c1df0ed9966a16c24add9ccb83/68747470733a2f2f646f63732e676f6f676c652e636f6d2f64726177696e67732f642f3132687034435071724e5046686174744c5f63496f4a707446766c41716d35774c513067677149356d6b43672f7075623f773d37343926683d333730 "Architecture")

## Service definition
[`pb/service.proto`](https://github.com/dray92/go-grpc-tutorial/blob/master/pb/service.proto)

## Golang Client/Server code generation
[`pb/service.pb.go`](https://github.com/dray92/go-grpc-tutorial/blob/master/pb/service.pb.go)

To generate:
```bash
$ protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=google/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
  pb/service.proto
```

### Golang Reverse proxy gateway code generation
[`pb/service.pb.gw.go`](https://github.com/dray92/go-grpc-tutorial/blob/master/pb/service.pb.gw.go)

To generate:
```bash
$ protoc -I/usr/local/include -I. \
   -I$GOPATH/src \
   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
   --grpc-gateway_out=logtostderr=true:. \
   pb/service.proto
```

### Golang Client code
[`client/client.go`](https://github.com/dray92/go-grpc-tutorial/blob/master/client/client.go)

### Golang Server code
[`server/server.go`](https://github.com/dray92/go-grpc-tutorial/blob/master/server/server.go)

### Golang Reverse proxy server code
[`server/server-rproxy.go`](https://github.com/dray92/go-grpc-tutorial/blob/master/server/server-rproxy.go)

### Golang Server <-> Client communication
```bash
go run server/server.go

...

go run client/client.go
```

Will display ``2018/05/28 12:57:11 Hello debo!``

### Golang Server <-> Client via gRPC Gateway
```bash
go run server/server.go

...

go run server/server-rproxy.go

...

curl -X POST "http://localhost:8080/v1/example/hello/v1/world | jq ."
```

Will display 
```json
{
  "id": "v1",
  "msg": "world"
}
```

## Python Client/Server code generation

PR to update gRPC docs: [#671](https://github.com/grpc/grpc.github.io/pull/671)
[`pb/service_pb2.py`](https://github.com/dray92/go-grpc-tutorial/blob/master/pb/service_pb2.py)
[`pb/service_pb2_grpc.py`](https://github.com/dray92/go-grpc-tutorial/blob/master/pb/service_pb2_grpc.py)

To generate:
```bash
virtualenv env
. env/bin/activate

pip install grpcio-tools googleapis-common-protos

python -m grpc_tools.protoc --python_out=. \
  --grpc_python_out=. -I. -I/usr/local/include \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/usr/local/include \
  pb/service.proto 
```

### Python Client code
[`client/client.py`](https://github.com/dray92/go-grpc-tutorial/blob/master/client/client.py)

### Python Server code
[`server/server.py`](https://github.com/dray92/go-grpc-tutorial/blob/master/server/server.py)

### Python Server <-> Client communication
```bash
python -m server.server

...

python -m client.client
```

Will display ``Hello world!``

### Python Reverse proxy server code
I don't believe this can be done in the state of the world at the time of writing.

## Generate Swagger JSON
[`pb/service.swagger.json`](https://github.com/dray92/go-grpc-tutorial/blob/master/pb/service.swagger.json)

To generate:
```bash
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:. \
  pb/service.proto
```