package mongodb

import (
	"context"

	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetMongoClient(uri string) (*mongo.Client, error) {
	clientCh := make(chan *mongo.Client)
	errorCh := make(chan error)

	go func() {
		client, err := mongo.Connect(options.Client().ApplyURI(uri))

		if err != nil {
			errorCh <- err
			return
		}

		if err := client.Ping(context.TODO(), nil); err != nil {
			errorCh <- err
			return
		}

		clientCh <- client
	}()

	select {
	case client := <-clientCh:
		return client, nil
	case err := <-errorCh:
		return nil, err
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
