package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Set up MongoDB connection parameters
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Create a MongoDB client
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Don't forget to defer closing the MongoDB client when you're done
	defer client.Disconnect(context.Background())

}
