package mongoRepositories

import (
	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	userCollection = "users"
)

func NewMongoDBUserRepository(db *mongo.Database) domain.UserRepository {
	return &mongoUserRepo{
		utils.NewMongoCrud[domain.User](db, userCollection),
	}
}

type mongoUserRepo struct {
	utils.MongoCrud[domain.User]
}
