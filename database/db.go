package database

import (
	"context"
	"time"

	"example.com/simplebank/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoDB().Host)); err != nil {

		return nil
	} else {
		return client
	}
}

var client = Connect()
