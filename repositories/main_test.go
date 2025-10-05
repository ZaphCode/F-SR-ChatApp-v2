package repositories

import (
	"context"
	"testing"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/lib/mongodb"
	"github.com/ZaphCode/F-SR-ChatApp/utils"
)

var conversationRepo domain.ConversationRepository
var messageRepo domain.MessageRepository
var userRepo domain.UserRepository

func TestMain(m *testing.M) {
	client := mongodb.MustGetMongoClient(utils.MONGO_DEV_URI)

	convColl := client.Database("test-db").Collection("conversations")
	userColl := client.Database("test-db").Collection("users")
	msgColl := client.Database("test-db").Collection("messages")

	conversationRepo = NewMongoDBConversationRepository(convColl)
	messageRepo = NewMongoDBMessageRepository(msgColl)
	userRepo = NewMongoDBUserRepository(userColl)

	m.Run()

	convColl.Drop(context.TODO())
	msgColl.Drop(context.TODO())
	userColl.Drop(context.TODO())

	client.Disconnect(context.TODO())
}
