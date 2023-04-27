GO_WORKSPACE := $(firstword $(subst ;, ,$(GOPATH)))\src
protoc:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. proto/*.proto

run:
	go run main.go