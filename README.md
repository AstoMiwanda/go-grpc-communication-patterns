# Communication Go gRPC Server and Client

It is based on the [gRPC Quickstart](https://grpc.io/docs/quickstart/go.html) and [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) tutorials.

There are communication have implemented:
- simple RPC
- server-side streaming RPC
- client-side streaming RPC
- bidirectional streaming RPC

# Running the application

1. Install the dependencies

```bash
go mod tidy
```

2. Run the server
`open new terminal`

```bash
cd server
go run *.go
```

3. Run the client
`open new terminal`

```bash
cd client
go run *.go
```