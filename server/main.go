package main

import (
	"github.com/astomiwanda/go-grpc-service/server/services/greet"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8080"
)

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Cannot listen %v", err)
	}

	server := grpc.NewServer()
	greet.InitGreetServer(server)
	log.Println("Server running on ", listen.Addr())
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Cannot serve grpc %v", err)
	}
}
