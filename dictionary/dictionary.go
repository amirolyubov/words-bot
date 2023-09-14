package dictionary

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"words-bot/db"
	"words-bot/gpt"
	"words-bot/users"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Translations struct {
	Ru string `bson:"ru,omitempty"`
	Fr string `bson:"fr,omitempty"`
}

type Meaning struct {
	Explanation  string `bson:"explanation" json:"explanation"`
	Example      string `bson:"example" json:"example"`
	PartOfSpeech string `bson:"part_of_speech" json:"part_of_speech"`
}

type Word struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Spelling      string             `bson:"spelling"`
	Meaning       []Meaning          `bson:"meaning"`
	Language      string             `bson:"language,omitempty"`
	Transcription string             `bson:"transcription,omitempty"`
	Translations  Translations       `bson:"translations,omitempty"`
	Synonyms      []string           `bson:"synonyms,omitempty"`
	Error         bool               `bson:"error,omitempty" json:"error,omitempty"`
}

func GetWord(spelling string) (Word, error) {
	collection, _ := db.GetCollection("words")
	filter := bson.M{"spelling": spelling}

	var word Word
	err := collection.FindOne(context.TODO(), filter).Decode(&word)

	if err != nil {
		return word, err
	}
	return word, nil
}

func GetWordById(wordId primitive.ObjectID) (Word, error) {
	collection, _ := db.GetCollection("words")
	filter := bson.M{"_id": wordId}

	var word Word
	err := collection.FindOne(context.TODO(), filter).Decode(&word)

	if err != nil {
		return word, err
	}
	return word, nil
}

func CreateNewWord(spelling string) (Word, error) {
	collection, _ := db.GetCollection("words")
	word := Word{}

	wordInfoString, err := gpt.GenerateWordInformation(spelling)
	if err != nil {
		return word, err
	}

	json.Unmarshal([]byte(wordInfoString), &word)

	if word.Error {
		return word, fmt.Errorf("error: there is no word")
	}

	result, err := collection.InsertOne(context.TODO(), word)

	if err != nil {
		return word, err
	}

	word.ID = result.InsertedID.(primitive.ObjectID)

	return word, nil

}

func AddWordToDictionary(wordId primitive.ObjectID, userId int64) error {
	collection, _ := db.GetCollection("users")

	filter := bson.D{{Key: "tg_id", Value: userId}}
	update := bson.M{
		"$addToSet": bson.M{"words": wordId},
	}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var user users.User

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return fmt.Errorf("Error adding word: ", err)
	}

	isWordAlreadyInDictfunc := CheckWordExistingInUserDictionary(wordId, userId)

	if isWordAlreadyInDictfunc {
		return fmt.Errorf("Error adding word: ", err)
	}

	err = collection.FindOneAndUpdate(
		context.TODO(),
		filter,
		update,
		options,
	).Decode(&user)

	if err != nil {
		return fmt.Errorf("Error adding word: ", err)
	}
	return nil
}

func CheckWordExistingInUserDictionary(wordId primitive.ObjectID, userId int64) bool {
	collection, _ := db.GetCollection("users")

	filter := bson.D{{Key: "tg_id", Value: userId}}

	var user users.User

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return false
	}

	for _, id := range user.Words {
		if id == wordId {
			return true
		}
	}
	return false
}

func GetRandomWords(total int) ([]Word, error) {
	collection, _ := db.GetCollection("words")
	pipeline := []bson.D{bson.D{{"$sample", bson.D{{"size", total}}}}}

	words := []Word{}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return words, err
	}
	for cursor.Next(context.TODO()) {
		var word Word
		err := cursor.Decode(&word)
		if err != nil {
			return words, err
		}
		words = append(words, word)
	}
	return words, nil
}

func GetRandomUserWord(userId int64) (Word, error) {
	collection, _ := db.GetCollection("users")
	filter := bson.D{{Key: "tg_id", Value: userId}}

	var user users.User
	var word Word

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return word, err
	}

	randWordId := user.Words[rand.Intn(len(user.Words))]
	word, err = GetWordById(randWordId)
	if err != nil {
		return word, err
	}

	return word, nil

}
