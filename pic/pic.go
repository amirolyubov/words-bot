package pic

import (
	"fmt"
	"log"
	"os"

	"github.com/fogleman/gg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateOneWordPic(word string, color string) tgbotapi.FileBytes {
	path := fmt.Sprintf(`./pic/%s.png`, word)
	const w = 1000
	const h = 350
	dc := gg.NewContext(w, h)
	switch color {
	case "white":
		dc.SetRGB(1, 1, 1)
	case "red":
		dc.SetRGB(1, 0.8, 0.8)
	case "green":
		dc.SetRGB(0.8, 1, 0.8)
	default:
		dc.SetRGB(1, 1, 1)

	}
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("./fonts/font.ttf", 100); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored(word, 40, h/2, 0, 0.5)
	dc.SavePNG(path)

	photoBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}

	err = os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}

	return photoFileBytes
}

func CreateWordExtendedPic(word string) tgbotapi.FileBytes {
	path := fmt.Sprintf(`./pic/%s.png`, word)
	const w = 1000
	const h = 1200
	dc := gg.NewContext(w, h)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("./fonts/font.ttf", 130); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored(word, 60, 150, 0, 0.5)
	dc.SavePNG(path)

	photoBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}

	err = os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}

	return photoFileBytes
}
