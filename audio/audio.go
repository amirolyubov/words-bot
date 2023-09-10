package audio

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	htgotts "github.com/hegedustibor/htgo-tts"
	handlers "github.com/hegedustibor/htgo-tts/handlers"
	voices "github.com/hegedustibor/htgo-tts/voices"
)

func CreateAudio(word string) (string, tgbotapi.FileBytes) {
	path := fmt.Sprintf("./audio/%s.mp3", word)
	speech := htgotts.Speech{Folder: "audio", Language: voices.EnglishUK, Handler: &handlers.MPlayer{}}
	speech.CreateSpeechFile(word, word)

	audioBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	audioFileBytes := tgbotapi.FileBytes{
		Name:  "voices",
		Bytes: audioBytes,
	}

	return path, audioFileBytes
}

func RemoveAudio(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}
}
