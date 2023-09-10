package messages

import (
	"fmt"
	"strings"
	"words-bot/audio"
	"words-bot/bot"
	"words-bot/pic"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BlankMessage(text string, chatId int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, text)

	return msg
}

func Card(word bot.Word, chatId int64) (tgbotapi.PhotoConfig, tgbotapi.VoiceConfig) {
	picPath, picBytes := pic.CreatePic(word.Spelling, word.Transcription)
	audioPath, audioBytes := audio.CreateAudio(word.Spelling)

	audioMsg := tgbotapi.NewVoice(chatId, audioBytes)

	picMsg := tgbotapi.NewPhoto(chatId, picBytes)
	picMsg.ParseMode = "Markdown"

	tgbotapi.NewPhoto(chatId, picBytes)

	head := fmt.Sprintf(`\[%s]

`, word.Transcription)
	meaning := ""
	for i, mean := range word.Meaning {
		meaning = meaning + fmt.Sprintf(`%s. %s
`,
			fmt.Sprint(i+1), mean.Explanation) + fmt.Sprintf(`_%#v_
`, mean.Example)
	}

	translate := fmt.Sprintf("\n_translation_\nru: %s\nfr: %s\n\n", word.Translations.Ru, word.Translations.Fr)

	synonyms := ""
	if len(word.Synonyms) > 0 {
		synonyms = fmt.Sprintf(`_synonyms_
%v`,
			strings.Join(word.Synonyms, ", "))
	}

	message := head + meaning + translate + synonyms

	picMsg.Caption = message

	pic.RemovePic(picPath)
	audio.RemoveAudio(audioPath)

	return picMsg, audioMsg
}
