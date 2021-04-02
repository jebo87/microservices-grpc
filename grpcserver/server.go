package grpcserver

import (
	"net"

	protobuffers "github.com/jebo87/microservices-grpc/protobuffers"
	"google.golang.org/grpc"
)

func StartServer() {

	grpcServer := grpc.NewServer()
	protobuffers.RegisterItemServiceServer(grpcServer, &ItemServer{})

	listener, err := net.Listen("tcp", "0.0.0.0:")

	if err != nil {
		panic(err)
	}
	defer listener.Close()

	grpcServer.Serve(listener)

}
