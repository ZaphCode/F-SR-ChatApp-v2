package utils

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoCrud[T any] struct {
	db         *mongo.Database
	collection string
}

func NewMongoCrud[T any](db *mongo.Database, collection string) MongoCrud[T] {
	return MongoCrud[T]{db, collection}
}

func (m *MongoCrud[T]) FindAll() ([]T, error) {
	cursor, err := m.db.Collection(m.collection).Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var docs []T

	if err = cursor.All(context.TODO(), &docs); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	return docs, nil
}

func (m *MongoCrud[T]) FindByID(id uuid.UUID) (T, error) {
	result := m.db.Collection(m.collection).FindOne(
		context.Background(), bson.D{{Key: "id", Value: id}},
	)

	var doc T

	if result.Err() != nil {
		return doc, result.Err()
	}

	if err := result.Decode(&doc); err != nil {
		return doc, err
	}

	return doc, nil
}

func (m *MongoCrud[T]) Save(doc *T) error {
	res, err := m.db.Collection(m.collection).InsertOne(context.TODO(), doc)

	if err != nil {
		return err
	}

	if res.InsertedID == nil {
		return errors.New("failed to insert data")
	}

	return nil
}

func (m *MongoCrud[T]) Update(id uuid.UUID, doc *T) error {
	_, err := m.db.Collection(m.collection).UpdateOne(
		context.Background(), bson.D{{Key: "id", Value: id}}, bson.D{{Key: "$set", Value: doc}},
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *MongoCrud[T]) Remove(id uuid.UUID) error {
	res, err := m.db.Collection(m.collection).DeleteOne(
		context.Background(), bson.D{{Key: "id", Value: id}},
	)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("resource not found")
	}

	return nil
}
