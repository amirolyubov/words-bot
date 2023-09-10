package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type GptMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GptRequestBody struct {
	Model       string       `json:"model"`
	Messages    []GptMessage `json:"messages"`
	Temperature float32      `json:"temperature"`
}

type GptResponseMessage struct {
	Content string `json:"content"`
}

type GptResponseAssistantRole struct {
	FinishReason string             `json:"finish_reason"`
	Index        int                `json:"index"`
	Message      GptResponseMessage `json:"message"`
	Role         string             `json:"role"`
}

type GptResponse struct {
	ID      string                     `json:"id"`
	Object  string                     `json:"object"`
	Created float64                    `json:"created"`
	Model   string                     `json:"model"`
	Choices []GptResponseAssistantRole `json:"choices"`
	Usage   map[string]int             `json:"usage"`
}

func GenerateWordInformation(word string) (string, error) {
	gptUrl := "https://api.openai.com/v1/chat/completions"

	data := GptRequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []GptMessage{
			{
				Role:    "user",
				Content: "найди данные для данного слова\n\nответ должен быть строго в формате json.\n ключи должны быть на аглийском языке в snakecase имена ключей следующие\n spelling,transcription, error, meaning (array of { explanation, example, part_of_speech } (если значение одно, должен быть массив с одним элементом, если значений несколько - тогда больше)),translations: {ru, fr}\n\nsynonims: (type array of string, max 3 length)\n\nесли ты не знаешь значения слова, пришли поле error: true\n\n" + word,
			},
		},
		Temperature: 0.7,
	}

	dataMarshalled, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, gptUrl, bytes.NewReader(dataMarshalled))
	if err != nil {
		log.Printf("Error: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("GPT_TOKEN"))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	defer res.Body.Close()

	var response GptResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
