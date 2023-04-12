package id_rpc

import (
	"context"
	"github.com/scaf-fold/srv-leaf-id/api"
)

type IdGenerator struct {
	api.UnimplementedIdServiceServer
}

func (id *IdGenerator) Id(context.Context, *api.IdRequest) (*api.IdResponse, error) {
	return nil, nil
}
func (id *IdGenerator) Ids(context.Context, *api.IdBatchRequest) (*api.IdBatchResponse, error) {
	return nil, nil
}
func (id *IdGenerator) Inverse(context.Context, *api.IdInverseRequest) (*api.IdInverseResponse, error) {
	return nil, nil
}
