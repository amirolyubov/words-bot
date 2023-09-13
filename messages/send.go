package messages

import (
	"words-bot/bot"
)

func Send(message string, chatId int64, lang string) {
	tgbot := bot.GetBot()

	word, err := bot.GetWord(message)
	if err != nil {
		word, err = bot.CreateNewWord(message)
		if err != nil {
			msg := BlankMessage("There is no word like this :|", chatId)
			tgbot.Send(msg)
			return
		}
	}

	isWordAlreadyInDict := bot.AddWordToDictionary(word.ID, chatId)

	card, audio := Card(word, chatId)
	if isWordAlreadyInDict != nil {
		card.Caption = card.Caption + "\n\n_already in your dict_"
	}
	tgbot.Send(card)
	tgbot.Send(audio)
}
