package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"order-service/internal/entities"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(client *MongoClient) *OrderRepository {
	return &OrderRepository{
		collection: client.DB.Collection("orders"),
	}
}

func (r *OrderRepository) Create(ctx context.Context, order *entities.Order) error {
	_, err := r.collection.InsertOne(ctx, order)
	return err
}

func (r *OrderRepository) GetByID(ctx context.Context, id string) (*entities.Order, error) {
	var order entities.Order
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) Update(ctx context.Context, order *entities.Order) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": order.ID}, bson.M{"$set": order})
	return err
}

func (r *OrderRepository) List(ctx context.Context, userID string) ([]entities.Order, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []entities.Order
	if err := cursor.All(ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}
