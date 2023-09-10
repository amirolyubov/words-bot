package pic

import (
	"fmt"
	"log"
	"os"

	"github.com/fogleman/gg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreatePic(word string, transcription string) (string, tgbotapi.FileBytes) {
	path := fmt.Sprintf(`./pic/%s.png`, word)
	const w = 1000
	const h = 350
	dc := gg.NewContext(w, h)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("./fonts/font.ttf", 100); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored(word, 30, h/2, 0, 0.5)
	dc.SavePNG(path)

	photoBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}

	return path, photoFileBytes
}

func RemovePic(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}
}
