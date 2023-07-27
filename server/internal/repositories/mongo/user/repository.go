package user

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	client *mongo.Client
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) List(ctx context.Context) {}
func (r *Repository) Create(ctx context.Context) {}