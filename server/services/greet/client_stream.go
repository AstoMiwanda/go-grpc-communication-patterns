package greet

import (
	"fmt"
	pb "github.com/astomiwanda/go-grpc-service/proto"
	"io"
	"log"
)

func (g *greetServer) SayHelloClientStreaming(stream pb.Greet_SayHelloClientStreamingServer) error {
	log.Println("Got stream for client")
	var messages []string
	for {
		receive, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		message := fmt.Sprintf("Hello, %s!", receive.Name)
		messages = append(messages, message)
	}
}
