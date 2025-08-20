package user

import (
	configgo "github.com/nextmillenniummedia/config-go"
)

type ClientConfig struct {
	Host string `config:"required,format=url"`
}

func GetConfig() (config ClientConfig, err error) {
	err = configgo.InitConfig(&config, configgo.Setting{
		Title:  "User client",
		Prefix: "USER",
	}).Process()
	return config, err
}
