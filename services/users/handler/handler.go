package handler

//type Users interface {
//	// MakeHat produces a hat of mysterious, randomly-selected color!
//	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
//}

import (
	"context"
	rpc2 "github.com/ostapetc/ms/tree/master/idgenerator/rpc"
	"github.com/ostapetc/ms/tree/master/users/rpc"
	//idgenerator "github.com/ostapetc/ms/tree/master/idgenerator/rpc"
	"github.com/ostapetc/ms/tree/master/users/service"
	"net/http"
)

type Handler struct {
	userService service.UsersService
}

func (h *Handler) Register(ctx context.Context, req *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	rpc2.NewIDGeneratorJSONClient("http://localhost:81", &http.Client{})
	//id :=

	//user := service.NewUser(id, req.Username, req.Password)
	//
	//token, err := h.userService.Register(user)
	//if err != nil {
	//
	//}

	client := idgenerator.NewIDGeneratorJSONClient("http://localhost:81", &http.Client{})
	res, err := client.GenerateID(context.Background(), &rpc.Request{Object: object})

	return &rpc.RegisterResponse{
		Id:    1,
		Token: "token",
	}, nil
}
