package mongodb

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
)

type mongoDBConversationRepo struct {
	mongoBaseCRUD[domain.Conversation]
}

func NewConversationRepository(coll *mongo.Collection) domain.ConversationRepository {
	return &mongoDBConversationRepo{
		newMongoBaseCRUD[domain.Conversation](coll),
	}
}

func (r *mongoDBConversationRepo) FindFrom(userA, userB uuid.UUID) (domain.Conversation, error) {
	filter := bson.D{{Key: "$or", Value: bson.A{
		bson.D{
			{Key: "usera.id", Value: userA},
			{Key: "userb.id", Value: userB},
		},
		bson.D{
			{Key: "usera.id", Value: userB},
			{Key: "userb.id", Value: userA},
		},
	}}}

	result := r.Coll.FindOne(context.TODO(), filter)

	var doc domain.Conversation

	if err := result.Err(); err != nil {
		return doc, err
	}

	if err := result.Decode(&doc); err != nil {
		return doc, err
	}

	return doc, nil
}

func (r *mongoDBConversationRepo) FindAllFrom(userID uuid.UUID) ([]domain.Conversation, error) {
	filter := bson.D{{Key: "$or", Value: bson.A{
		bson.D{{Key: "usera.id", Value: userID}},
		bson.D{{Key: "userb.id", Value: userID}},
	}}}

	result, err := r.Coll.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	var docs []domain.Conversation

	if err := result.All(context.TODO(), &docs); err != nil {
		return nil, err
	}

	return docs, nil
}
