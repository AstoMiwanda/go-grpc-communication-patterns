package greet

import (
	"context"
	pb "github.com/astomiwanda/go-grpc-service/proto"
	"io"
	"log"
)

func CallSayHelloServerStreaming(client pb.GreetClient) {
	names := []string{"Lorem", "Ipsum", "Dolor", "Amet"}
	streaming, err := client.SayHelloServerStreaming(context.Background(), &pb.NamesList{Names: names})
	if err != nil {
		log.Fatalf("Error client %v", err)
	}
	for {
		receive, err := streaming.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receive %v", err)
		}
		log.Println(receive.Message)
	}

}
