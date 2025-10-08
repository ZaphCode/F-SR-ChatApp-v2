package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/utils"
)

type mongoDBUserRepo struct {
	mongoBaseCRUD[domain.User]
}

func NewUserRepository(coll *mongo.Collection) domain.UserRepository {
	return &mongoDBUserRepo{
		newMongoBaseCRUD[domain.User](coll),
	}
}

func (r *mongoDBUserRepo) FindByEmail(email string) (domain.User, error) {
	result := r.Coll.FindOne(
		context.Background(), bson.D{{Key: "email", Value: email}},
	)

	var doc domain.User

	if err := result.Err(); err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return doc, err
		}

		return doc, utils.ErrNotFound
	}

	if err := result.Decode(&doc); err != nil {
		return doc, err
	}

	return doc, nil
}
