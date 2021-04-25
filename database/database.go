package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connecting to the database
func ConnectDB() *mongo.Collection {
	fmt.Println("Initilizing Connection")
	options := options.Client().ApplyURI(os.Getenv("MONGOATLASURL"))

	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the Database")

	return client.Database("books").Collection("book")
}
