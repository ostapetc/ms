package handler

import (
	"context"
	idgRpc "ms/services/idg/rpc"
	"ms/services/users/rpc"
	"ms/services/users/service"
	"net/http"
)

type Handler struct {
	service  *service.UsersService
	idgCient idgRpc.IDGenerator
}

func NewHandler(idgServiceAddr string) *Handler {
	storage := service.NewStorage()
	validator := service.RegisterValidator{}

	handler := &Handler{}
	handler.service = service.NewUsersService(storage, validator)
	handler.idgCient = idgRpc.NewIDGeneratorJSONClient(idgServiceAddr, &http.Client{})

	return handler
}

func (h *Handler) Register(ctx context.Context, req *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	hasher := service.UserHasher{}

	id, err := h.getUserID()
	if err != nil {
		return &rpc.RegisterResponse{Error: err.Error()}, nil
	}

	user := service.NewUser(id, req.Username, req.Password, hasher)

	token, err := h.service.Register(user)
	if err != nil {
		return &rpc.RegisterResponse{Error: err.Error()}, nil
	}

	return &rpc.RegisterResponse{
		Id:    id,
		Token: token,
	}, nil
}

func (h Handler) getUserID() (int64, error) {
	req := &idgRpc.Request{Object: "users"}
	res, err := h.idgCient.GenerateID(context.Background(), req)

	if err != nil {
		return 0, err
	}

	return res.Id, nil
}
