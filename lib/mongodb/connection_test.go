package mongodb

import (
	"context"
	"testing"
)

func TestGetMongoClient(t *testing.T) {
	client, err := GetMongoClient("mongodb://zaph:zaphpass@localhost:27017")

	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	_, err = GetMongoClient("mongo://localhosttt:270171")

	if err == nil {
		t.Fatalf("Expected error for invalid MongoDB URI, got nil")
	}
}
