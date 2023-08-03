package user

import (
	"context"

	"github.com/mieltn/keepintouch/internal/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	client *mongo.Client
}

type dbUser struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*dto.User, error) {
	collection := r.client.Database("keepintouch").Collection("users")

	var dbItem dbUser
	if err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&dbItem); err != nil {
		return nil, err
	}

	return &dto.User{
		Id: dbItem.Id.Hex(),
		Username: dbItem.Username,
		Email: dbItem.Email,
		Password: dbItem.Password,
	}, nil
}

func (r *Repository) Create(ctx context.Context, in *dto.UserCreateReq) (*dto.User, error) {
	collection := r.client.Database("keepintouch").Collection("users")
	res, err := collection.InsertOne(
		ctx,
		bson.D{
			{"username", in.Username},
			{"email", in.Email},
			{"password", in.Password},
		},
	)
	if err != nil {
		return nil, err
	}
	return &dto.User{
		Id:       res.InsertedID.(primitive.ObjectID).Hex(),
		Username: in.Username,
		Email:    in.Email,
	}, nil
}
