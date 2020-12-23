package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewClient returns a mongodb client based on passed connection string.
func NewClient(ctx context.Context, cs string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cs))
	if err != nil {
		return nil, fmt.Errorf("Database connection error: %s", err)
	}

	return client, nil
}

// NewDB returns a mongodb database based on passed client and name.
func NewDB(ctx context.Context, client *mongo.Client, name string) (*mongo.Database, error) {
	database := client.Database(name)
	err := client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("Connection to database failed: %s", err)
	}

	return database, nil
}
