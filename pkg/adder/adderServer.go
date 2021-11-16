package adder

import (
	"context"
	"grpctest/pkg/api"
)

type AdderServer struct {
	api.UnimplementedAdderServer
}

func (s *AdderServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
