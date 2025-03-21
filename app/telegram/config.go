package telegram

import (
	configgo "github.com/nextmillenniummedia/config-go"
)

type TelegramConfig struct {
	Token   string `config:"required"`
	Timeout int    `config:"default=60"`
	Debug   bool   `config:""`
}

func GetConfig() (config TelegramConfig, err error) {
	err = configgo.InitConfig(&config, configgo.Setting{
		Title:  "Telegram",
		Prefix: "TELEGRAM",
	}).Process()
	return config, err
}
