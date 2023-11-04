package message

import (
	"context"
	"time"

	"github.com/mieltn/keepintouch/internal/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

type dbMessage struct {
	Id        primitive.ObjectID `bson:"_id"`
	ChatId    primitive.ObjectID `bson:"chat_id"`
	UserId    primitive.ObjectID `bson:"user_id"`
	Text      string             `bson:"text"`
	CreatedAt primitive.DateTime `bson:"created_at"`
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (m dbMessage) ToDTO() *dto.Message {
	return &dto.Message{
		Id:        m.Id.Hex(),
		ChatId:    m.ChatId.Hex(),
		UserId:    m.UserId.Hex(),
		Text:      m.Text,
		CreatedAt: m.CreatedAt.Time(),
	}
}

func (r *Repository) MessageByChatId(ctx context.Context, in *dto.MessageByChatIdReq) ([]*dto.Message, error) {
	collection := r.client.Database("keepintouch").Collection("messages")
	filter := bson.M{}

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}
	filter["chat_id"] = oid

	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})
	opts.SetLimit(int64(in.Limit))
	opts.SetSkip(int64(in.Offset))

	cur, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var dbItems []*dbMessage
	if err := cur.All(ctx, &dbItems); err != nil {
		return nil, err
	}

	messages := make([]*dto.Message, len(dbItems))
	for i, item := range dbItems {
		messages[i] = item.ToDTO()
	}

	return messages, nil
}

func (r *Repository) Create(ctx context.Context, in *dto.MessageCreateReq) (*dto.Message, error) {
	collection := r.client.Database("keepintouch").Collection("messages")

	chatOid, err := primitive.ObjectIDFromHex(in.ChatId)
	if err != nil {
		return nil, err
	}

	userOid, err := primitive.ObjectIDFromHex(in.UserId)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	res, err := collection.InsertOne(
		ctx,
		bson.D{
			{"chat_id", chatOid},
			{"user_id", userOid},
			{"text", in.Text},
			{"created_at", primitive.NewDateTimeFromTime(now)},
		},
	)
	if err != nil {
		return nil, err
	}
	return &dto.Message{
		Id:       res.InsertedID.(primitive.ObjectID).Hex(),
		ChatId:     in.ChatId,
		UserId: in.UserId,
		Text: in.Text,
		CreatedAt: now,
	}, nil
}
