package mongodb

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
)

// Example struct to be used with MongoCrud
type toDo struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreateAt    time.Time `json:"createAt"`
}

var mongoCrud MongoCrud[toDo]

func TestMain(m *testing.M) {
	//* On startup
	client, err := GetMongoClient("mongodb://zaph:zaphpass@localhost:27017")

	if err != nil {
		panic(err)
	}

	db := client.Database("test-db")

	mongoCrud = NewMongoCrud[toDo](db.Collection("todos"))

	m.Run()

	//* On shutdown
	// if err := db.Drop(context.Background()); err != nil {
	// 	panic(err)
	// }

	if err := client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

func TestMongoCrud_Create(t *testing.T) {
	todo := toDo{
		ID:          uuid.New(),
		Description: "Test todo",
		Completed:   false,
		CreateAt:    time.Now(),
	}

	if err := mongoCrud.Save(&todo); err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}
}

func TestMongoCrud_FindAll(t *testing.T) {
	todos, err := mongoCrud.FindAll()

	if err != nil {
		t.Fatalf("Failed to find todos: %v", err)
	}

	for _, todo := range todos {
		todoJSON, err := json.MarshalIndent(todo, "", "  ")
		if err != nil {
			t.Fatalf("Failed to marshal todo: %v", err)
		}
		t.Logf("Todo: %s", todoJSON)
	}
}
