package app

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"os/signal"
	"syscall"

	"gate/clients/user"
	. "gate/errors"
	"gate/telegram"
	"gate/types"
	"gate/utils/log"

	loggergo "github.com/nextmillenniummedia/logger-go"
	"google.golang.org/grpc"
)

type app struct {
	logger   loggergo.ILogger
	configs  Configs
	telegram types.ITelegram
	clients  struct {
		user user.UserClient
	}
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

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	userConnection, err := grpc.NewClient("localhost:15000", opts...)
	WriteErrorAndExit(err, a.logger)
	a.clients.user = user.NewUserClient(userConnection)
	userAddData := user.AddUserRequest{
		Name: "Eugen",
	}
	response, err := a.clients.user.AddUser(context.Background(), &userAddData)
	fmt.Println("response", response, "err", err)

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
