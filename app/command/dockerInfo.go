package command

import (
	"deli13/dockerbot/app/docker"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

type DockerInfo struct {
	Command string
}

func (c DockerInfo) Has(message tgbotapi.Update) bool {
	if message.CallbackQuery != nil {
		log.Info().Msg("Обработчик не определён")
		if strings.Contains(message.CallbackQuery.Data, c.Command) {
			log.Info().Str("Callback", message.CallbackQuery.Data).Msg("Обработчик совпал")
			return true
		}
	}
	log.Info().Msg("Обработчик не определён")
	return false
}

func (c DockerInfo) Work(bot *tgbotapi.BotAPI, upate tgbotapi.Update) {
	log.Info().Msg("Обрабатываем callback сообщение")
	containerId := strings.Replace(upate.CallbackQuery.Data, c.Command, "", 1)
	infoContainer := docker.ContainerInfo(containerId)
	message := "ID контецнера: " + infoContainer.Id + "\nИмя контейнера: " + infoContainer.Name + "\nОбраз: " + infoContainer.Image + "\nСтатус: " + infoContainer.Status + "\n Статистика:\n" + infoContainer.Stats
	msg := tgbotapi.NewMessage(upate.CallbackQuery.From.ID, message)
	var keyboard []tgbotapi.InlineKeyboardButton
	keyboard = append(keyboard, tgbotapi.NewInlineKeyboardButtonData("start", "start_"+infoContainer.Id))
	keyboard = append(keyboard, tgbotapi.NewInlineKeyboardButtonData("restart", "restart_"+infoContainer.Id))
	keyboard = append(keyboard, tgbotapi.NewInlineKeyboardButtonData("stop", "stop_"+infoContainer.Id))
	fmt.Println(keyboard)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(keyboard...))
	bot.Send(msg)
	bot.Send(tgbotapi.NewCallback(upate.CallbackQuery.ID, "Выполнено"))

}
