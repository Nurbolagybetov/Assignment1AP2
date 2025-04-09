package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoClient struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongoClient(opts *Options) (*MongoClient, error) {
	clientOptions := options.Client().ApplyURI(opts.URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the database to ensure connection
	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	db := client.Database(opts.Database)
	return &MongoClient{
		Client: client,
		DB:     db,
	}, nil
}

func (m *MongoClient) Disconnect() {
	if err := m.Client.Disconnect(context.Background()); err != nil {
		log.Printf("Error disconnecting MongoDB: %v", err)
	}
}
