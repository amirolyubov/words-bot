package schedule

import (
	"words-bot/messages"

	cron "github.com/robfig/cron/v3"
)

func InitSchedule() {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("0 0 11,16,20 * * *", func() { messages.SendRandomWord() })
	// c.AddFunc("*/4 * * * * *", messages.SendRandomWord)

	c.Start()
}
