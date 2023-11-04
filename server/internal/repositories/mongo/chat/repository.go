package chat

import (
	"context"

	"github.com/mieltn/keepintouch/internal/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

type dbChat struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Password string             `bson:"password"`
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (m dbChat) ToDTO() *dto.Chat {
	return &dto.Chat{
		Id:       m.Id.Hex(),
		Name:     m.Name,
	}
}

func (r *Repository) List(ctx context.Context, in *dto.ChatListReq) ([]*dto.Chat, error) {
	collection := r.client.Database("keepintouch").Collection("chats")
	filter := bson.M{}

	if len(in.Ids) > 0 {
		oids := make([]primitive.ObjectID, len(in.Ids))
		for i := range in.Ids {
			objId, err := primitive.ObjectIDFromHex(in.Ids[i])
			if err != nil {
				return nil, err
			}
			oids[i] = objId
		}
		filter["_id"] = bson.M{"$in": oids}
	}

	opts := options.Find()
	opts.SetLimit(int64(in.Limit))
	opts.SetSkip(int64(in.Offset))

	cur, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var dbItems []*dbChat
	if err := cur.All(ctx, &dbItems); err != nil {
		return nil, err
	}

	chats := make([]*dto.Chat, len(dbItems))
	for i, item := range dbItems {
		chats[i] = item.ToDTO()
	}

	return chats, nil
}

func (r *Repository) Create(ctx context.Context, in *dto.ChatCreateReq) (*dto.Chat, error) {
	collection := r.client.Database("keepintouch").Collection("chats")

	res, err := collection.InsertOne(
		ctx,
		bson.D{{"name", in.Name}, {"password", in.Password}},
	)
	if err != nil {
		return nil, err
	}
	return &dto.Chat{
		Id:       res.InsertedID.(primitive.ObjectID).Hex(),
		Name:     in.Name,
	}, nil
}
