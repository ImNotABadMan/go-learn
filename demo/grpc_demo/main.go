package main

import (
	"context"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcPeopleServer struct {
	UnimplementedPeopleServer
}

func main() {
	grpcServer := grpc.NewServer()
	RegisterPeopleServer(grpcServer, new(GrpcPeopleServer))

	listener, err := net.Listen("tcp", "192.168.10.113:9998")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
	defer grpcServer.Stop()

}

func (h *GrpcPeopleServer) SayHello(context.Context, *Hello) (*Hello, error) {
	hello := Hello{
		Name: "Server",
		Text: "answer hello",
	}

	_, err := proto.Marshal(&hello)
	if err != nil {
		log.Fatal(err)
	}

	return &hello, err
}
