package db

import (
	"context"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

func NewMongoDB(ctx context.Context) (*MongoDB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	return &MongoDB{
		client: client,
	}, nil
}

func (d *MongoDB) Close(ctx context.Context) error {
	if err := d.client.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}

func (d *MongoDB) GetDB() *mongo.Client {
	return d.client
}

