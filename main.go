package main

import (
	"log"
	"os"
	"words-bot/bot"
	"words-bot/db"
	"words-bot/dictionary"
	"words-bot/messages"
	"words-bot/schedule"
	"words-bot/users"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	db.InitDb(os.Getenv("DB_URI"))
	schedule.InitSchedule()

	err := <-bot.Init(os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	tgbot := bot.GetBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tgbot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ParseMode = "Markdown"
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.IsCommand() {
				if update.Message.Text == "/start" {
					err := users.CreateNewUser(update.Message.From.ID, update.Message.From.UserName)
					if err != nil {
						msg.Text = "you are already here"
					} else {
						msg.Text = "success. you are here"
					}

				}
				tgbot.Send(msg)
			} else {
				tgbot := bot.GetBot()

				word, err := dictionary.GetWord(update.Message.Text)
				if err != nil {
					word, err = dictionary.CreateNewWord(update.Message.Text)
					if err != nil {
						msg := messages.BlankMessage("There is no word like this :|", update.Message.From.ID)
						tgbot.Send(msg)
						return
					}
				}

				isWordAlreadyInDict := dictionary.AddWordToDictionary(word.ID, update.Message.From.ID)

				card, audio := messages.Card(word, update.Message.From.ID)
				if isWordAlreadyInDict != nil {
					card.Caption = card.Caption + "\n\n_already in your dict_"
				}
				tgbot.Send(card)
				tgbot.Send(audio)
			}
		}
	}
}
