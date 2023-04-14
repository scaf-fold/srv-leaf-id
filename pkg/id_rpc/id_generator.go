package id_rpc

import (
	"context"
	"errors"
	"github.com/scaf-fold/srv-leaf-id/api"
	"github.com/scaf-fold/srv-leaf-id/global"
	"strconv"
)

type IdGenerator struct {
	api.UnimplementedIdServiceServer
}

func (id *IdGenerator) Id(context.Context, *api.IdRequest) (*api.IdResponse, error) {
	ids, err := global.ID.Ids(1)
	if err != nil {
		return nil, err
	}
	if ids != nil && len(ids) > 0 {
		return &api.IdResponse{
			Id: ids[0],
		}, nil
	}
	return nil, errors.New("StatusInternalServerError")
}
func (id *IdGenerator) Ids(context.Context, *api.IdBatchRequest) (*api.IdBatchResponse, error) {
	ids, err := global.ID.Ids(1)
	if err != nil {
		return nil, err
	}
	return &api.IdBatchResponse{
		Ids: ids,
	}, nil
}
func (id *IdGenerator) Inverse(_ context.Context, r *api.IdInverseRequest) (*api.IdInverseResponse, error) {
	snow, err := global.ID.IdInverse(r.Id, global.StartTime)
	if err != nil {
		return nil, err
	}
	return &api.IdInverseResponse{
		Id:       snow.Id,
		Time:     snow.GenerateTime.String(),
		Sequence: strconv.FormatInt(snow.Sequence, 10),
		WorkId:   strconv.FormatInt(snow.WorkId, 10),
	}, nil
}
