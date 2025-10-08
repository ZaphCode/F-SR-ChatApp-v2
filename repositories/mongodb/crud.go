package mongodb

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/ZaphCode/F-SR-ChatApp/utils"
)

type mongoBaseCRUD[T any] struct {
	Coll *mongo.Collection
}

func newMongoBaseCRUD[T any](collection *mongo.Collection) mongoBaseCRUD[T] {
	return mongoBaseCRUD[T]{collection}
}

func (m *mongoBaseCRUD[T]) FindAll() ([]T, error) {
	cursor, err := m.Coll.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	var docs []T

	err = cursor.All(context.TODO(), &docs)

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	return docs, nil
}

func (m *mongoBaseCRUD[T]) FindByID(id uuid.UUID) (T, error) {
	result := m.Coll.FindOne(
		context.TODO(), bson.D{{Key: "id", Value: id}},
	)

	var doc T

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

func (m *mongoBaseCRUD[T]) Save(doc *T) error {
	res, err := m.Coll.InsertOne(context.TODO(), doc)

	if err != nil {
		return err
	}

	if res.InsertedID == nil {
		return errors.New("failed to insert data")
	}

	return nil
}

func (m *mongoBaseCRUD[T]) Update(id uuid.UUID, doc *T) error {
	res, err := m.Coll.UpdateOne(
		context.TODO(), bson.D{{Key: "id", Value: id}}, bson.D{{Key: "$set", Value: doc}},
	)

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return utils.ErrNotFound
	}

	return nil
}

func (m *mongoBaseCRUD[T]) Remove(id uuid.UUID) error {
	res, err := m.Coll.DeleteOne(
		context.TODO(), bson.D{{Key: "id", Value: id}},
	)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return utils.ErrNotFound
	}

	return nil
}
