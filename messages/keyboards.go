package messages

import (
	"fmt"
	"words-bot/dictionary"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func KeyboardWithRandomWords(words []dictionary.Word, corrent primitive.ObjectID) tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup()

	for _, word := range words {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(word.Translations.Ru, fmt.Sprintf("quiz,%s,%s", word.ID.Hex(), corrent.Hex())),
		))
	}

	return keyboard
}

func KeyboardNextQuizOrNo() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next", "next,yes"),
			tgbotapi.NewInlineKeyboardButtonData("Stop", "next,no"),
		),
	)

	return keyboard
}
