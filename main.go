package main

import (
	"deli13/dockerbot/app/bot"
	"deli13/dockerbot/app/config"

	"github.com/rs/zerolog/log"
)

func main() {
	config.LoadConfig("project.yaml")
	log.Info().Msg("Конфигурация загружена")
	bot.StartBot(config.Config)
}
