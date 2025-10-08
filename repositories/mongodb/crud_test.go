package mongodb

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/utils"
)

func TestMongoCrud_Save(t *testing.T) {
	todo := toDo{
		ID:          uuid.New(),
		Description: "Test todo",
		Completed:   false,
		CreateAt:    time.Now(),
	}

	if err := mongoCrud.Save(&todo); err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	if err := mongoCrud.Save(nil); err == nil {
		t.Fatalf("Expected error when saving nil todo, got nil")
	}
}

func TestMongoCrud_FindAll(t *testing.T) {
	if err := mongoCrud.Save(&toDo{
		ID:          uuid.New(),
		Description: "Test todo 1",
		Completed:   false,
		CreateAt:    time.Now(),
	}); err != nil {
		t.Errorf("Failed to create todo: %v", err)
	}

	if err := mongoCrud.Save(&toDo{
		ID:          uuid.New(),
		Description: "Test todo 2",
		Completed:   true,
		CreateAt:    time.Now(),
	}); err != nil {
		t.Errorf("Failed to create todo: %v", err)
	}

	todos, err := mongoCrud.FindAll()

	if err != nil {
		t.Fatalf("Failed to find todos: %v", err)
	}

	if len(todos) != 2 {
		t.Fatalf("Expected at least two todos, got %d", len(todos))
	}

	utils.PrettyPrint(todos)
}

func TestMongoCrud_FindByID(t *testing.T) {
	todo := toDo{
		ID:          uuid.New(),
		Description: "Test todo",
		Completed:   false,
		CreateAt:    time.Now(),
	}

	if err := mongoCrud.Save(&todo); err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	foundTodo, err := mongoCrud.FindByID(todo.ID)

	if err != nil {
		t.Errorf("Failed to find todo by ID: %v", err)
	}

	if foundTodo.ID != todo.ID {
		t.Errorf("Expected todo ID %v, got %v", todo.ID, foundTodo.ID)
	}

	if _, err := mongoCrud.FindByID(uuid.New()); err == nil {
		t.Errorf("Expected error when finding todo with non-existent ID, got nil")
	}

	if _, err := mongoCrud.FindByID(uuid.Nil); err == nil {
		t.Errorf("Expected error when finding todo with nil ID, got nil")
	}
}

func TestMongoCrud_UpdateProperly(t *testing.T) {
	todo := toDo{
		ID:          uuid.New(),
		Description: "Test todo",
		Completed:   false,
		CreateAt:    time.Now(),
	}

	if err := mongoCrud.Save(&todo); err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	// Original todo
	utils.PrettyPrint(todo)

	todo.Description = "Updated todo"
	todo.Completed = true

	if err := mongoCrud.Update(todo.ID, &todo); err != nil {
		t.Fatalf("Failed to update todo: %v", err)
	}

	foundTodo, err := mongoCrud.FindByID(todo.ID)

	if err != nil {
		t.Errorf("Failed to find updated todo by ID: %v", err)
	}

	if foundTodo.Description != "Updated todo" {
		t.Errorf("Expected updated todo description 'Updated todo', got '%s'", foundTodo.Description)
	}

	// Updated todo
	utils.PrettyPrint(foundTodo)
}

func TestMongoCrud_UpdateBad(t *testing.T) {
	todo := toDo{
		ID:          uuid.New(),
		Description: "Test todo",
		Completed:   false,
		CreateAt:    time.Now(),
	}

	if err := mongoCrud.Save(&todo); err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	if err := mongoCrud.Update(uuid.New(), &todo); err == nil && !errors.Is(err, utils.ErrNotFound) {
		t.Errorf("Expected error when updating non-existent todo, got nil")
	}

	if err := mongoCrud.Update(todo.ID, nil); err == nil {
		t.Errorf("Expected error when updating with nil todo, got nil")
	}

	if err := mongoCrud.Update(uuid.Nil, &todo); err == nil {
		t.Errorf("Expected error when updating with nil ID, got nil")
	}
}

func TestMongoCrud_Remove(t *testing.T) {
	todo := toDo{
		ID:          uuid.New(),
		Description: "Test todo",
		Completed:   false,
		CreateAt:    time.Now(),
	}

	if err := mongoCrud.Save(&todo); err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	if err := mongoCrud.Remove(todo.ID); err != nil {
		t.Fatalf("Failed to delete todo: %v", err)
	}

	if _, err := mongoCrud.FindByID(todo.ID); err == nil {
		t.Errorf("Expected error when finding deleted todo, got nil")
	}

	if err := mongoCrud.Remove(uuid.New()); !errors.Is(err, utils.ErrNotFound) {
		t.Errorf("Expected error when removing non-existent todo, got nil")
	}

	if err := mongoCrud.Remove(uuid.Nil); err == nil {
		t.Errorf("Expected error when removing todo with nil ID, got nil")
	}
}
