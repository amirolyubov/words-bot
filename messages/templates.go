package messages

import (
	"log"
	"math/rand"
	"words-bot/audio"
	"words-bot/dictionary"
	"words-bot/games"
	"words-bot/pic"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BlankMessage(text string, chatId int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, text)

	return msg
}

func Card(word dictionary.Word, chatId int64) (tgbotapi.PhotoConfig, tgbotapi.VoiceConfig) {
	picBytes := pic.CreateOneWordPic(word.Spelling, "white")
	audioBytes := audio.CreateAudio(word.Spelling)

	audioMsg := tgbotapi.NewVoice(chatId, audioBytes)

	picMsg := tgbotapi.NewPhoto(chatId, picBytes)
	picMsg.ParseMode = "Markdown"

	tgbotapi.NewPhoto(chatId, picBytes)

	message := CardCaption(word)

	picMsg.Caption = message

	return picMsg, audioMsg
}

func CardWithActions(word dictionary.Word, chatId int64) tgbotapi.PhotoConfig {
	randomWords, err := dictionary.GetRandomWords(2)
	if err != nil {
		log.Fatal(err)
	}
	randomWords = append(randomWords, word)

	rand.Shuffle(len(randomWords), func(i, j int) { randomWords[i], randomWords[j] = randomWords[j], randomWords[i] })

	picBytes := pic.CreateOneWordPic(word.Spelling, "white")

	msg := tgbotapi.NewPhoto(chatId, picBytes)
	msg.ParseMode = "Markdown"

	message := QuizCaption(word)

	msg.Caption = message

	msg.ReplyMarkup = KeyboardWithRandomWords(randomWords, word.ID)
	msg.DisableNotification = false

	tgbotapi.NewPhoto(chatId, picBytes)

	return msg
}

func UpdateQuizCard(chatId int64, messageId int, result games.QuizResult) (tgbotapi.EditMessageMediaConfig, tgbotapi.EditMessageCaptionConfig) {
	continueKeyboard := KeyboardNextQuizOrNo()
	var newPic tgbotapi.FileBytes
	if result.Result {
		newPic = pic.CreateOneWordPic(result.Correct.Translations.Ru, "green")
	} else {
		newPic = pic.CreateOneWordPic(result.Correct.Translations.Ru, "red")
	}
	newWordPic := tgbotapi.EditMessageMediaConfig{
		BaseEdit: tgbotapi.BaseEdit{
			MessageID:   messageId,
			ChatID:      chatId,
			ReplyMarkup: &continueKeyboard,
		},
		Media: tgbotapi.NewInputMediaPhoto(newPic),
	}

	message := QuizCaption(result.Correct)

	newWordCaption := tgbotapi.NewEditMessageCaption(chatId, messageId, message)
	newWordCaption.ReplyMarkup = &continueKeyboard
	newWordCaption.ParseMode = "Markdown"

	return newWordPic, newWordCaption
}
