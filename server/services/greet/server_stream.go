package greet

import (
	"fmt"
	pb "github.com/astomiwanda/go-grpc-service/proto"
	"log"
	"time"
)

func (g *greetServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.Greet_SayHelloServerStreamingServer) error {
	log.Printf("Got request with names: %v", req.Names)
	for _, name := range req.Names {
		message := fmt.Sprintf("Hello, %s!", name)
		if err := stream.Send(&pb.HelloResponse{Message: message}); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
