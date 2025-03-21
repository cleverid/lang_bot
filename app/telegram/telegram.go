package telegram

import (
	"lb/app/types"
	"lb/app/utils/log"

	loggergo "github.com/nextmillenniummedia/logger-go"
)

type telegram struct {
	config TelegramConfig
	logger loggergo.ILogger
}

func New(config TelegramConfig, logger loggergo.ILogger) (tel types.ITelegram, err error) {
	tel = &telegram{
		config: config,
		logger: logger.Clone().From(log.FROM_TELEGRAM),
	}
	return tel, nil
}

func (t *telegram) Start() {
	t.logger.Info("started")
}

func (t *telegram) Stop() error {
	t.logger.Info("stopped")
	return nil
}
