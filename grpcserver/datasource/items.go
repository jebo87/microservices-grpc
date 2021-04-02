package datasource

import (
	"errors"

	protobuffers "github.com/jebo87/microservices-grpc/protobuffers"
)

var (
	list = &protobuffers.Items{
		Items: []*protobuffers.Item{
			{Id: 1, Name: "Item 1", Description: "This is my item 1", Price: 150},
			{Id: 2, Name: "Item 2", Description: "This is my item 2", Price: 10.5},
			{Id: 3, Name: "Item 3", Description: "This is my item 3", Price: 99},
			{Id: 4, Name: "Item 4", Description: "This is my item 4", Price: 20},
			{Id: 5, Name: "Item 5", Description: "This is my item 5", Price: 36.5},
		},
	}
)

func GetAll() (*protobuffers.Items, error) {
	return list, nil
}

func GetSingle(id int64) (*protobuffers.Item, error) {
	for _, a := range list.Items {
		if a.Id == id {
			return a, nil
		}
	}
	return nil, errors.New("not found")
}
