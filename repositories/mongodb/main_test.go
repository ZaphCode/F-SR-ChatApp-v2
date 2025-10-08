package mongodb

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/utils"
)

type toDo struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreateAt    time.Time `json:"createAt"`
}

var mongoCrud mongoBaseCRUD[toDo]
var conversationRepo domain.ConversationRepository
var messageRepo domain.MessageRepository
var userRepo domain.UserRepository

func TestMain(m *testing.M) {
	utils.PrettyPrint("Setting up test database...")

	client := MustGetMongoClient(utils.MONGO_DEV_URI)
	db := client.Database(utils.MONGO_DEV_DB)

	todosColl := client.Database("test-db").Collection("todos")
	convColl := db.Collection(utils.MONGO_CONVERSATION_COL)
	userColl := db.Collection(utils.MONGO_USER_COL)
	msgColl := db.Collection(utils.MONGO_MESSAGE_COL)

	mongoCrud = newMongoBaseCRUD[toDo](todosColl)
	conversationRepo = NewConversationRepository(convColl)
	messageRepo = NewMessageRepository(msgColl)
	userRepo = NewUserRepository(userColl)

	m.Run()

	utils.PrettyPrint("Tearing down test database...")

	convColl.Drop(context.TODO())
	msgColl.Drop(context.TODO())
	userColl.Drop(context.TODO())

	client.Disconnect(context.TODO())
}
