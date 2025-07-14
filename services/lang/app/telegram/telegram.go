package telegram

import (
	"lb/services/lang/app/types"
	"lb/services/lang/app/utils/log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	loggergo "github.com/nextmillenniummedia/logger-go"
)

type telegram struct {
	config TelegramConfig
	logger loggergo.ILogger
}

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

var numericKeyboard2 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

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
				// msg.ReplyToMessageID = update.Message.MessageID
				switch update.Message.Command() {
				case "translate":
					msg.ReplyMarkup = numericKeyboard2
				}
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				// msg.ReplyMarkup = numericKeyboard
				bot.Send(msg)
			} else if update.CallbackQuery != nil {
				// Respond to the callback query, telling Telegram to show the user
				// a message with the data received.
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
				if _, err := bot.Request(callback); err != nil {
					panic(err)
				}

				// // And finally, send a message containing the data received.
				// msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
				// if _, err := bot.Send(msg); err != nil {
				// 	panic(err)
				// }
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
