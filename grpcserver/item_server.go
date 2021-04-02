package grpcserver

import (
	"context"

	"github.com/jebo87/microservices-grpc/grpcserver/datasource"
	protobuffers "github.com/jebo87/microservices-grpc/protobuffers"
)

type ItemServer struct{}

func (ItemServer) List(ctx context.Context, void *protobuffers.Void) (*protobuffers.Items, error) {
	return datasource.GetAll()
}
func (ItemServer) ItemInfo(ctx context.Context, itemId *protobuffers.ItemID) (*protobuffers.Item, error) {
	return datasource.GetSingle(itemId.Id)
}
