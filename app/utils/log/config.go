package log

import (
	configgo "github.com/nextmillenniummedia/config-go"
)

type LoggerConfig struct {
	Level    string  `config:"enum=verbose|debug|info|warn|error|fatal|silent,default=error,doc='Log level'"`
	Pretty   bool    `config:"doc='Switch to pretty mod'"`
	Sampling float64 `config:"doc='Percent of sampling in range from 0.0 to 1.0'"`
}

func GetConfig() (config LoggerConfig) {
	configgo.InitConfig(&config, configgo.Setting{
		Title:  "Logger",
		Prefix: "LOG",
	}).Process()
	return
}
