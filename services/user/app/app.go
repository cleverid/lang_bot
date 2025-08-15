package app

import (
	"context"
	"errors"
	"fmt"
	"user/contracts"
)

type app struct {
	contracts.UnimplementedUserServer
}

func New() contracts.UserServer {
	return &app{}
}

func (a *app) AddUser(ctx context.Context, request *contracts.AddUserRequest) (*contracts.AddUserResponse, error) {
	fmt.Println(request)
	response := contracts.AddUserResponse{
		UserId: "user_id",
		Name:   request.Name,
		Address: &contracts.Address{
			Street: "street",
		},
	}
	err := errors.New("test error")
    err = nil
	return &response, err
}
