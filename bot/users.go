package bot

import (
	"context"
	"fmt"
	"words-bot/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID   `bson:"_id,omitempty"`
	TgID  int64                `bson:"tg_id,omitempty"`
	Name  string               `bson:"name,omitempty"`
	Words []primitive.ObjectID `bson:"words"`
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
		TgID:  userId,
		Name:  name,
		Words: make([]primitive.ObjectID, 0),
	}
	// Insert user document into MongoDB
	insertingResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return fmt.Errorf("Error creating user: ", err)
	}

	fmt.Println(insertingResult)

	return nil
}
