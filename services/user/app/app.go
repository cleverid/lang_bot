package app

import (
	"context"
	"user/contracts"
)

type app struct {
	contracts.UnimplementedUserServer
}

func New() contracts.UserServer {
	return &app{}
}

func (a *app) AddUser(context.Context, *contracts.AddUserRequest) (*contracts.AddUserResponse, error) {
	return nil, nil
}
