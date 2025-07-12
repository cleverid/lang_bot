package app

import "lb/contracts"

func NewGRPCServer() {
	response := contracts.AddUserResponse{}
	contracts.RegisterUserServer()
}
