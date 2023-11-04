package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (sp *serviceProvider) GetDB() *mongo.Client {
	if sp.db == nil {
		client, err := mongo.Connect(sp.ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			sp.addError(err)
			return nil
		}
		sp.onClose(client.Disconnect)
		sp.db = client
	}
	return sp.db
}