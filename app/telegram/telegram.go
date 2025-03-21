package telegram

import (
	"lb/app/types"
	"lb/app/utils/log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

func (t *telegram) Start() error {
	bot, err := tgbotapi.NewBotAPI(t.config.Token)
	if err != nil {
		return err
	}
	bot.Debug = t.config.Debug

	go func() {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = t.config.Timeout
		updates := bot.GetUpdatesChan(u)
		for update := range updates {
			if update.Message != nil { // If we got a message
				t.logger.Info("Message",
					"user_name", update.Message.From.UserName,
					"text", update.Message.Text)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	}()

	t.logger.Info("started", "user", bot.Self.UserName)
	return nil
}

func (t *telegram) Stop() error {
	t.logger.Info("stopped")
	return nil
}
