package greet

import (
	"fmt"
	pb "github.com/astomiwanda/go-grpc-service/proto"
	"io"
	"log"
	"time"
)

func (g *greetServer) SayHelloBidirectionalStreaming(stream pb.Greet_SayHelloBidirectionalStreamingServer) error {
	log.Println("Got bidirectional stream")
	for {
		receive, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println("Got request name", receive.Name)
		message := fmt.Sprintf("Hello, %s!", receive.Name)
		if err = stream.Send(&pb.HelloResponse{Message: message}); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
}
