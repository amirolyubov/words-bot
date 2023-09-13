package users

import (
	"context"
	"fmt"
	"words-bot/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	TgID           int64                `bson:"tg_id,omitempty"`
	Name           string               `bson:"name,omitempty"`
	ActiveLanguage string               `bson:"active_languange,omitempty" json:"active_languange,omitempty"`
	Words          []primitive.ObjectID `bson:"words"`
}

func CreateNewUser(userId int64, name string) error {
	collection, _ := db.GetCollection("users")
	filter := bson.M{"tg_id": userId}
	var existingUser User

	err := collection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err == nil {
		return fmt.Errorf("you are already here")
	}

	user := User{
		TgID:           userId,
		Name:           name,
		Words:          make([]primitive.ObjectID, 0),
		ActiveLanguage: "ru",
	}
	// Insert user document into MongoDB
	insertingResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return fmt.Errorf("Error creating user: ", err)
	}

	fmt.Println(insertingResult)

	return nil
}

func GetMe(userId int64) (User, error) {
	collection, _ := db.GetCollection("users")
	filter := bson.M{"tg_id": userId}
	var user User

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err == nil {
		return user, err
	}

	return user, nil
}
