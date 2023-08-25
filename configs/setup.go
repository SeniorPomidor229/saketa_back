package configs

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

func NewDB() *DB {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("employee_directory")
	collection := database.Collection("employees")

	return &DB{
		Client:     client,
		Database:   database,
		Collection: collection,
	}
}

func (db *DB) Close() {
	db.Client.Disconnect(context.Background())
}