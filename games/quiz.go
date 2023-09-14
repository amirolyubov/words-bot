package games

import (
	"words-bot/dictionary"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuizResult struct {
	Result  bool
	Correct dictionary.Word
	Attempt dictionary.Word
}

func ProcessQuizResult(attempt string, correct string) (QuizResult, error) {
	result := QuizResult{
		Result: attempt == correct,
	}

	attemptId, err := primitive.ObjectIDFromHex(attempt)
	if err != nil {
		return result, err
	}
	correctId, err := primitive.ObjectIDFromHex(correct)
	if err != nil {
		return result, err
	}

	attemptWord, err := dictionary.GetWordById(attemptId)
	if err != nil {
		return result, err
	}
	correctWord, err := dictionary.GetWordById(correctId)
	if err != nil {
		return result, err
	}

	result.Attempt = attemptWord
	result.Correct = correctWord

	return result, nil
}
