package audio

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateAudio(word string) tgbotapi.FileBytes {
	requestURL := fmt.Sprintf("http://0.0.0.0:59125/api/tts?text=%s&voice=en_UK/apope_low&noiseScale=0.667&noiseW=0.8&lengthScale=1&ssml=false", url.QueryEscape(word))
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	audioFileBytes := tgbotapi.FileBytes{
		Name:  "voices",
		Bytes: body,
	}

	return audioFileBytes
}
