package greet

import (
	"context"
	pb "github.com/astomiwanda/go-grpc-service/proto"
	"log"
)

func CallSayHello(client pb.GreetClient) {
	data, err := client.SayHello(context.Background(), &pb.NoParam{})
	if err != nil {
		log.Fatalf("Error client %v", err)
	}
	log.Println(data.Message)
}
