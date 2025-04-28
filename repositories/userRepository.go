package repositories

import (
	"context"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/lib/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type mongoDBUserRepo struct {
	mongodb.MongoCrud[domain.User]
}

func NewMongoDBUserRepository(coll *mongo.Collection) domain.UserRepository {
	return &mongoDBUserRepo{
		mongodb.NewMongoCrud[domain.User](coll),
	}
}

func (r *mongoDBUserRepo) FindByEmail(email string) (domain.User, error) {
	result := r.Coll.FindOne(
		context.Background(), bson.D{{Key: "email", Value: email}},
	)

	var doc domain.User

	if err := result.Err(); err != nil {
		return doc, err
	}

	if err := result.Decode(&doc); err != nil {
		return doc, err
	}

	return doc, nil
}
