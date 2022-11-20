package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartCommand struct {
	Command         string
	AllowedKeyboard []string
}

func (c StartCommand) Has(message *tgbotapi.Message) bool {
	if message.Text == c.Command {
		return true
	}
	return false
}

func (c StartCommand) Work(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	bot.Send(tgbotapi.NewSetMyCommands(tgbotapi.BotCommand{
		Command:     c.Command,
		Description: "Запуск",
	}))
	var keyboard []tgbotapi.KeyboardButton
	for _, str := range c.AllowedKeyboard {
		keyboard = append(keyboard, tgbotapi.NewKeyboardButton(str))
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "Доступные действия")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(keyboard...))
	bot.Send(msg)
}
