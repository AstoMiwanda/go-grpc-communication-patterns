package greet

import (
	"context"
	pb "github.com/astomiwanda/go-grpc-service/proto"
	"io"
	"log"
	"time"
)

func SayHelloBidirectionalStreaming(client pb.GreetClient) {
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error client %v", err)
	}

	wait := make(chan struct{})
	go func() {
		for {
			receive, err2 := stream.Recv()
			if err2 == io.EOF {
				break
			}
			if err2 != nil {
				log.Fatalf("Error receive %v", err2)
			}
			log.Println(receive.Message)
		}
		close(wait)
	}()

	names := []string{"Lorem", "Ipsum", "Dolor", "Amet"}
	for _, name := range names {
		log.Println("Send name", name)
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalf("Error send %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Error close send %v", err)
	}
	<-wait
	log.Println("Bi stream is done")
}
