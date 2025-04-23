package mongodb

import (
	"context"

	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type result struct {
	Client *mongo.Client
	Err    error
}

func GetMongoClient(uri string) (*mongo.Client, error) {
	resultChan := make(chan result)

	go func() {
		client, err := mongo.Connect(options.Client().ApplyURI(uri))

		if err != nil {
			resultChan <- result{nil, err}
			return
		}

		if err := client.Ping(context.TODO(), nil); err != nil {
			resultChan <- result{nil, err}
			return
		}
		resultChan <- result{client, nil}
	}()

	select {
	case res := <-resultChan:
		return res.Client, res.Err
	case <-time.After(3 * time.Second):
		return nil, errors.New("timeout: conexión a MongoDB tardó demasiado")
	}
}

func MustGetMongoClient(uri string) *mongo.Client {
	client, err := GetMongoClient(uri)

	if err != nil {
		panic(err)
	}

	return client
}
