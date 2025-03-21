package telegram

import (
	configgo "github.com/nextmillenniummedia/config-go"
)

type TelegramConfig struct {
	Token string `config:"required"`
}

func GetConfig() (config TelegramConfig, err error) {
	err = configgo.InitConfig(&config, configgo.Setting{
		Title:  "Telegram",
		Prefix: "TELEGRAM",
	}).Process()
	return config, err
}
