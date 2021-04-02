package main

import "github.com/jebo87/microservices-grpc/grpcserver"

func main() {
	chann := make(chan bool)

	go grpcserver.StartServer()

	<-chann
}
