package command

import (
	"deli13/dockerbot/app/docker"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DockerList struct {
	Command string
}

func (c DockerList) Has(message *tgbotapi.Message) bool {
	if message.Text == c.Command {
		return true
	}
	return false
}

func (c DockerList) Work(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	listContainer := docker.ListContainer()
	var keyboard []tgbotapi.InlineKeyboardButton
	for _, cont := range listContainer {
		keyboard = append(keyboard, tgbotapi.NewInlineKeyboardButtonData(cont, cont))
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "Список контейнеров:")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(keyboard)
	bot.Send(msg)
}
