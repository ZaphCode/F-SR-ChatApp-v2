package repositories

import (
	"context"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/lib/mongodb"
	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type mongoDBConversationRepo struct {
	mongodb.MongoCrud[domain.Conversation]
}

func NewMongoDBConversationRepository(coll *mongo.Collection) domain.ConversationRepository {
	return &mongoDBConversationRepo{
		mongodb.NewMongoCrud[domain.Conversation](coll),
	}
}

func (r *mongoDBConversationRepo) FindFrom(userA, userB uuid.UUID) (domain.Conversation, error) {
	filter := bson.D{{Key: "$or", Value: bson.A{
		bson.D{
			{Key: "userid_a", Value: userA},
			{Key: "userid_b", Value: userB},
		},
		bson.D{
			{Key: "userid_a", Value: userB},
			{Key: "userid_b", Value: userA},
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
		bson.D{{Key: "userid_a", Value: userID}},
		bson.D{{Key: "userid_b", Value: userID}},
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
