package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx = func() context.Context {
	return context.Background()
}()

func Connect() (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	err = client.Connect(Ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("belajar_golang"), nil
}
