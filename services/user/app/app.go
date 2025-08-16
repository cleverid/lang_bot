package app

import (
	"context"
	"errors"
	"fmt"
	"user/clients/user"
)

type app struct {
	user.UnimplementedUserServer
}

func New() user.UserServer {
	return &app{}
}

func (a *app) AddUser(ctx context.Context, request *user.AddUserRequest) (*user.AddUserResponse, error) {
	fmt.Println(request)
	response := user.AddUserResponse{
		UserId: "user_id",
		Name:   request.Name,
		Address: &user.Address{
			Street: "street",
		},
	}
	err := errors.New("test error")
	err = nil
	return &response, err
}
