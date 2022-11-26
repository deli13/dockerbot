package command

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type CommandInterface interface {
	Has(message tgbotapi.Update) bool
	Work(bot *tgbotapi.BotAPI, message tgbotapi.Update)
}
