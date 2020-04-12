package handler

import (
	"context"
	idgRpc "ms/services/idg/rpc"
	"ms/services/users/rpc"
	"ms/services/users/service"
	"net/http"
)

type Handler struct {
	service *service.UsersService
}

func NewHandler() *Handler {
	storage := service.NewStorage()
	validator := service.RegisterValidator{}

	handler := &Handler{}
	handler.service = service.NewUsersService(storage, validator)

	return handler
}

func (h *Handler) Register(ctx context.Context, req *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	hasher := service.UserHasher{}

	id, err := getUserID("users")
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

//todo: pass server addr
func getUserID(object string) (int64, error) {
	client := idgRpc.NewIDGeneratorJSONClient("http://localhost:80", &http.Client{})

	req := &idgRpc.Request{Object: object}
	res, err := client.GenerateID(context.Background(), req)

	if err != nil {
		return 0, err
	}

	return res.Id, nil
}
