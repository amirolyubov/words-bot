package main

import (
	"log"
	"strings"
	"words-bot/bot"
	"words-bot/db"
	"words-bot/dictionary"
	"words-bot/games"
	"words-bot/messages"
	"words-bot/schedule"
	"words-bot/users"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	db.InitDb()
	schedule.InitSchedule()

	err := <-bot.Init()
	if err != nil {
		panic(err)
	}

	tgbot := bot.GetBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tgbot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			callback.Text = ""
			tgbot.Request(callback)
			callbackData := strings.Split(update.CallbackQuery.Data, ",")

			switch callbackData[0] {
			case "quiz":
				result, _ := games.ProcessQuizResult(callbackData[1], callbackData[2])

				newWordPic, newWordCaption := messages.UpdateQuizCard(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, result)

				tgbot.Request(newWordPic)
				tgbot.Request(newWordCaption)
			case "next":
				switch callbackData[1] {
				case "yes":
					word, _ := dictionary.GetRandomUserWord(update.SentFrom().ID)
					msg := messages.CardWithActions(word, update.CallbackQuery.From.ID)

					deleteMessageRequest := tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)

					tgbot.Request(deleteMessageRequest)
					tgbot.Send(msg)

				case "no":
					deleteMessageRequest := tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
					tgbot.Request(deleteMessageRequest)
				default:
					continue
				}
			default:
				continue
			}

		} else if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
					err := users.CreateNewUser(update.Message.From.ID, update.Message.From.UserName)
					if err != nil {
						msg.Text = "you are already here"
					} else {
						msg.Text = "success. you are here"
					}
					tgbot.Send(msg)
				case "quiz":
					word, _ := dictionary.GetRandomUserWord(update.SentFrom().ID)
					msg := messages.CardWithActions(word, update.Message.From.ID)

					tgbot.Send(msg)
				default:
					continue
				}
			} else {
				tgbot := bot.GetBot()

				word, err := dictionary.GetWord(update.Message.Text)
				if err != nil {
					word, err = dictionary.CreateNewWord(update.Message.Text)
					if err != nil {
						msg := messages.BlankMessage("There is no word like this :|", update.Message.From.ID)
						tgbot.Send(msg)
						continue
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
