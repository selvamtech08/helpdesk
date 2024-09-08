package store

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DatabaseName = "helpdesk"
)

var (
	DB *mongo.Client = ConnectDB()
)

func ConnectDB() *mongo.Client {
	mongoOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("mongo db connected")
	client.Database("helpdesk")
	return client
}

func GetCollection(collection string) *mongo.Collection {
	return DB.Database(DatabaseName).Collection(collection)
}
