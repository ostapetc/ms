package handler

import (
	"context"
	"github.com/ostapetc/ms/tree/master/idgenerator/rpc"
	"github.com/ostapetc/ms/tree/master/idgenerator/service"
)

type Handler struct {
	generator service.IDGenerator
}

func (h *Handler) GenerateID(ctx context.Context, req *rpc.Request) (*rpc.Response, error) {
	return &rpc.Response{
		Id: h.generator.Generate(req.Object),
	}, nil
}
