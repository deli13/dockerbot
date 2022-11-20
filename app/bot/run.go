package bot

import (
	"deli13/dockerbot/app/command"
	"deli13/dockerbot/app/config"
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

var botInstance *tgbotapi.BotAPI

func Instance(config config.Configuration) *tgbotapi.BotAPI {
	var err error
	if botInstance == nil {
		botInstance, err = tgbotapi.NewBotAPI(config.Credentials.TgCredentials)
		if err != nil {
			panic(err)
		}
		log.Info().Msg("Получено соединение с telegram")
	}
	return botInstance
}

func StartBot(config config.Configuration) {
	botApi := Instance(config)
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 10
	botApi.Debug = true
	messages := botApi.GetUpdatesChan(update)
	for message := range messages {
		if message.Message == nil {
			continue
		}
		err := authHandler(config, *message.Message)
		if err != nil {
			log.Warn().
				Err(err).
				Str("имя пользователя", message.Message.From.UserName).
				Int64("ID пользователя", message.Message.From.ID).
				Str("Текст", message.Message.Text).
				Msg("Попытка не авторизованного доступа")
			errorResponse := tgbotapi.NewMessage(message.Message.Chat.ID, err.Error())
			errorResponse.ReplyToMessageID = message.Message.MessageID
			botApi.Send(errorResponse)
			continue
		}
		proc := Processor()
		for _, c := range proc {
			if c.Has(message.Message) {
				c.Work(botApi, message.Message)
			}
		}

	}
}

func authHandler(config config.Configuration, msg tgbotapi.Message) error {
	for _, user := range config.Project.AllowUser {
		if user == msg.Chat.ID {
			return nil
		}
	}
	return errors.New("Access Denied")
}

func Processor() []command.CommandInterface {
	handlers := []command.CommandInterface{}
	start := command.StartCommand{AllowedKeyboard: []string{"Docker List"}, Command: "/start"}
	list := command.DockerList{Command: "Docker List"}
	handlers = append(handlers, start, list)
	return handlers
}
