package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	. "lb/app/errors"
	"lb/app/telegram"
	"lb/app/types"
	"lb/app/utils"
	"lb/app/utils/log"

	loggergo "github.com/nextmillenniummedia/logger-go"
)

type app struct {
	logger   loggergo.ILogger
	configs  Configs
	telegram types.ITelegram
}

func Init() *app {
	utils.LoadEnv(".env")
	app := &app{}
	logger := log.New(log.GetConfig())

	configs, err := getConfigs()
	WriteErrorAndExit(err, logger)
	app.configs = configs

	telegram, err := telegram.New(configs.Telegram, logger)
	WriteErrorAndExit(err, logger)
	app.telegram = telegram

	app.logger = logger

	return app
}

func (a *app) Start() {
	a.telegram.Start()

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
