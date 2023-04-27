package greet

import (
	pb "github.com/astomiwanda/go-grpc-service/proto"
	"google.golang.org/grpc"
)

type greetServer struct {
	pb.UnimplementedGreetServer
}

func InitGreetServer(server *grpc.Server) {
	pb.RegisterGreetServer(server, &greetServer{})
}
