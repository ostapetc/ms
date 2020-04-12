package handler

import (
	"context"
	log "github.com/sirupsen/logrus"
	"ms/services/idg/rpc"
	"ms/services/idg/service"
)

type Handler struct {
	service service.IDGenerator
}

func (h *Handler) GenerateID(ctx context.Context, req *rpc.Request) (*rpc.Response, error) {
	id := h.service.Generate(req.Object)

	log.Infof("Generated ID %s:%d\n", req.Object, id)

	return &rpc.Response{
		Id: id,
	}, nil
}
