package main

import (
	"fmt"
	"github.com/astomiwanda/go-grpc-service/client/services/greet"
	pb "github.com/astomiwanda/go-grpc-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	port = ":8080"
)

func main() {
	host := fmt.Sprintf("localhost%s", port)
	dial, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error dial %v", err)
	}
	defer func() {
		err := dial.Close()
		if err != nil {
			log.Fatalf("Error close dial %v", err)
		}
	}()

	client := pb.NewGreetClient(dial)
	//greet.CallSayHello(client)
	//greet.CallSayHelloServerStreaming(client)
	//greet.CallSayHelloClientStreaming(client)
	greet.SayHelloBidirectionalStreaming(client)
}
