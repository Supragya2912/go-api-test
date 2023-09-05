package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(client *mongo.Client, user User) error {

	collection := client.Database("userDb").Collection("users")

	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func GetUsers(client *mongo.Client, userID string) (User, error) {

	collection := client.Database("userDb").Collection("users")

	var user User
	filter := bson.M{"_id": userID}
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	return user, err
}
