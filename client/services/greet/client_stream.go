package greet

import (
	"context"
	pb "github.com/astomiwanda/go-grpc-service/proto"
	"log"
	"time"
)

func CallSayHelloClientStreaming(client pb.GreetClient) {
	names := []string{"Lorem", "Ipsum", "Dolor", "Amet"}
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error client %v", err)
	}

	for _, name := range names {
		log.Println("Send name", name)
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalf("Error send %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	receive, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receive %v", err)
	}
	log.Println("Done Stream Client")
	log.Println(receive.Messages)
}
