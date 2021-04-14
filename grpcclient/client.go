package grpcclient

import (
	"context"
	"encoding/json"
	"log"

	protobuffers "github.com/jebo87/microservices-grpc/protobuffers"
	"google.golang.org/grpc"
)

func StartClient() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := protobuffers.NewItemServiceClient(conn)

	log.Println("Returning the complete list")
	result, err := client.List(context.Background(), &protobuffers.Void{})
	if err != nil {
		panic(err)
	}
	jsonList, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		panic(err)
	}
	log.Println(string(jsonList))

	log.Println("Returning item 3")
	itemID := &protobuffers.ItemID{Id: 3}
	item, err := client.ItemInfo(context.Background(), itemID)
	if err != nil {
		panic(err)
	}
	jsonItem, err := json.MarshalIndent(item, "", "\t")
	if err != nil {
		panic(err)
	}
	log.Println(string(jsonItem))

}
