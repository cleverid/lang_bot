package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	. "lb/services/lang/app/errors"
	"lb/services/lang/app/telegram"
	"lb/services/lang/app/types"
	"lb/services/lang/app/utils/log"

	loggergo "github.com/nextmillenniummedia/logger-go"
)

type app struct {
	logger   loggergo.ILogger
	configs  Configs
	telegram types.ITelegram
}

func Init() *app {
	app := &app{}

	logger := log.New(log.GetConfig())
	app.logger = logger

	configs, err := getConfigs()
	WriteErrorAndExit(err, logger)
	app.configs = configs

	telegram, err := telegram.New(configs.Telegram, logger)
	WriteErrorAndExit(err, logger)
	app.telegram = telegram

	return app
}

func (a *app) Start() {
	err := a.telegram.Start()
	WriteErrorAndExit(err, a.logger)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c
	fmt.Println() // For new line before stopping
	errs := a.Stop()
	WriteErrorsAndExit(errs, a.logger)
}

func (a *app) Stop() []error {
	errors := make([]error, 0)
	return errors
}
