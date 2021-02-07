package main

import (
	"context"
	"demo/grpc_demo/hello"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcPeopleServer struct {
	hello.PeopleServer
}

func main() {
	grpcServer := grpc.NewServer()
	hello.RegisterPeopleServer(grpcServer, new(GrpcPeopleServer))

	listener, err := net.Listen("tcp", "192.168.10.113:9998")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Start Server 192.168.10.113:9998")
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
	defer grpcServer.Stop()

}

func (h *GrpcPeopleServer) SayHello(context.Context, *hello.Hello) (*hello.Hello, error) {
	pb := hello.Hello{
		Name: "Server",
		Text: "answer hello",
	}

	_, err := proto.Marshal(&pb)
	fmt.Println("Send", pb.Name)
	if err != nil {
		log.Fatal(err)
	}

	return &pb, err
}
func (h *GrpcPeopleServer) SayHelloAgain(ctx context.Context, inHello *hello.Hello) (*hello.Hello, error) {
	return &hello.Hello{
		Name: "Server 192.168.10.113:9998",
		Text: "Say Hello Again",
	}, nil
}
