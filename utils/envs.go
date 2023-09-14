package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type Envs struct {
	BotToken string
	GptToken string
	DbUri    string
}

func GetEnvs() (Envs, error) {
	godotenv.Load()
	var envs Envs

	envs.BotToken = os.Getenv("BOT_TOKEN")
	envs.GptToken = os.Getenv("GPT_TOKEN")
	envs.DbUri = os.Getenv("DB_URI")

	return envs, nil
}
