package configs

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(EnvMongoURI()))

	if err != nil {
		log.Fatal(err)
	}

	// Ping to database
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Print(collections)
	fmt.Println("Connected to MongoDB")

	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// Getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	db := client.Database("digital_bikes")
	collection := db.Collection(collectionName)
	return collection
}
