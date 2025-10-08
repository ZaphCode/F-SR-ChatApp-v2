package mongodb

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
)

type mongoDBMessageRepo struct {
	mongoBaseCRUD[domain.Message]
}

func NewMessageRepository(coll *mongo.Collection) domain.MessageRepository {
	return &mongoDBMessageRepo{
		newMongoBaseCRUD[domain.Message](coll),
	}
}

func (r *mongoDBMessageRepo) FindAllFrom(conv uuid.UUID) ([]domain.Message, error) {
	filter := bson.D{{Key: "conversationid", Value: conv}}

	result, err := r.Coll.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	var messages []domain.Message

	if err := result.All(context.Background(), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
