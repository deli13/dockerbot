package command

import (
	"deli13/dockerbot/app/docker"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

type DockerHandler struct {
	StartCommand, StopCommand, RestartCommand string
}

func (c DockerHandler) Has(message tgbotapi.Update) bool {
	if message.CallbackQuery != nil {
		log.Info().Msg("Обработчик не определён")
		if strings.Contains(message.CallbackQuery.Data, c.StartCommand) || strings.Contains(message.CallbackQuery.Data, c.StopCommand) || strings.Contains(message.CallbackQuery.Data, c.RestartCommand) {
			log.Info().Str("Callback", message.CallbackQuery.Data).Msg("Обработчик совпал")
			return true
		}
	}
	log.Info().Msg("Обработчик не определён")
	return false
}

func (c DockerHandler) Work(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Info().Msg("Обработываем сообщение")
	command, containerID := c.matchCommand(update.CallbackQuery.Data)
	var err error
	if command == c.StartCommand {
		err = docker.StartContainer(containerID)
	}
	if command == c.StopCommand {
		err = docker.StopContainer(containerID)
	}
	if command == c.RestartCommand {
		err = docker.RestartContainer(containerID)
	}
	if err != nil {
		bot.Send(tgbotapi.NewCallback(update.CallbackQuery.ID, "Произошла ошибка при выполнение команды"))
		return
	}
	bot.Send(tgbotapi.NewCallback(update.CallbackQuery.ID, "Выполнена команда: "+command+", контейнер: "+containerID))

}

func (c DockerHandler) matchCommand(data string) (string, string) {
	if strings.Contains(data, c.RestartCommand) {
		return c.RestartCommand, strings.Replace(data, c.RestartCommand, "", 1)
	}
	if strings.Contains(data, c.StartCommand) {
		return c.StartCommand, strings.Replace(data, c.StartCommand, "", 1)
	}
	if strings.Contains(data, c.StopCommand) {
		return c.StopCommand, strings.Replace(data, c.StopCommand, "", 1)
	}
	return "", ""
}
