package greet

import (
	"context"
	pb "github.com/astomiwanda/go-grpc-service/proto"
)

func (g *greetServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello"}, nil
}
