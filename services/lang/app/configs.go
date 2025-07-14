package app

import (
	"errors"
	"fmt"
	"lb/services/lang/app/telegram"
	"lb/services/lang/app/utils/log"
)

type Configs struct {
	Logger   log.LoggerConfig
	Telegram telegram.TelegramConfig
}

func getConfigs() (configs Configs, err error) {
	errs := make([]error, 0)

	logger := log.GetConfig()

	telegram, err := telegram.GetConfig()
	errs = append(errs, err)

	configs = Configs{
		Logger:   logger,
		Telegram: telegram,
	}
	err = errors.Join(errs...)
	if err != nil {
		err = fmt.Errorf("configuration errors: \n%w", err)
	}
	return configs, err
}
