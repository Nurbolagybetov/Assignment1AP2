package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"inventory-service/internal/entities"
	"log"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(client *MongoClient) *ProductRepository {
	return &ProductRepository{
		collection: client.DB.Collection("products"),
	}
}

func (r *ProductRepository) Create(ctx context.Context, product *entities.Product) error {
	log.Printf("Attempting to insert product: %+v", product)
	result, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		log.Printf("Failed to insert product: %v", err)
		return err
	}
	log.Printf("Successfully inserted product with ID: %v", result.InsertedID)
	return nil
}

func (r *ProductRepository) GetByID(ctx context.Context, id string) (*entities.Product, error) {
	var product entities.Product
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Update(ctx context.Context, product *entities.Product) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": product.ID}, bson.M{"$set": product})
	return err
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *ProductRepository) List(ctx context.Context, offset, limit int) ([]entities.Product, error) {
	findOptions := options.Find()
	findOptions.SetSkip(int64(offset))
	findOptions.SetLimit(int64(limit))

	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []entities.Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}
