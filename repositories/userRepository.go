package repositories

import (
	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/lib/mongodb"
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
