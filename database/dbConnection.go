package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDb := "mongodb://localhost:27017"
	fmt.Println("Connecting to:", MongoDb)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(MongoDb)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Mongo Connect Error:", err)
	}

	// Ping to confirm connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB Ping Failed:", err)
	}

	fmt.Println("✅ Connected to MongoDB")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("restaurant").Collection(collectionName)
}
