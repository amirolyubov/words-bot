package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot tgbotapi.BotAPI

func Init(token string) <-chan error {
	c := make(chan error)

	go func() {
		newBot, err := tgbotapi.NewBotAPI(token)

		if !bot.Self.IsBot {
			c <- err
		}

		bot = *newBot

		c <- nil
	}()

	return c
}

func GetBot() *tgbotapi.BotAPI {
	return &bot
}
