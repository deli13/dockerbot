package command

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type CommandInterface interface {
	Has(message *tgbotapi.Message) bool
	Work(bot *tgbotapi.BotAPI, message *tgbotapi.Message)
}
