package command

import (
	"deli13/dockerbot/app/docker"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DockerList struct {
	Command string
}

func (c DockerList) Has(message tgbotapi.Update) bool {
	if message.Message != nil && message.Message.Text == c.Command {
		return true
	}
	return false
}

func (c DockerList) Work(bot *tgbotapi.BotAPI, message tgbotapi.Update) {
	listContainer := docker.ListContainer()
	var keyboard []tgbotapi.InlineKeyboardButton
	for _, cont := range listContainer {
		keyboard = append(keyboard, tgbotapi.NewInlineKeyboardButtonData(cont.Name, "info_"+cont.Id))
	}
	msg := tgbotapi.NewMessage(message.Message.Chat.ID, "Список контейнеров:")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(makeSliceChunk(keyboard, 3)...)
	bot.Send(msg)
}

func makeSliceChunk(keyboard []tgbotapi.InlineKeyboardButton, sizeChunk int) [][]tgbotapi.InlineKeyboardButton {
	var inlineSlice [][]tgbotapi.InlineKeyboardButton
	size := len(keyboard) / (len(keyboard) / sizeChunk)
	var first, last int
	for i := 0; i < len(keyboard)/size+1; i++ {
		first = i * size
		last = first + size
		if last > len(keyboard) {
			last = len(keyboard)
		}
		if first == last {
			break
		}
		inlineSlice = append(inlineSlice, keyboard[first:last])
	}
	return inlineSlice
}
