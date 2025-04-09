package usecase

import (
	"context"
	"inventory-service/internal/entities"
)

type ProductRepository interface {
	Create(ctx context.Context, product *entities.Product) error
	GetByID(ctx context.Context, id string) (*entities.Product, error)
	Update(ctx context.Context, product *entities.Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]entities.Product, error)
}

type ProductUsecase interface {
	CreateProduct(ctx context.Context, product *entities.Product) error
	GetProduct(ctx context.Context, id string) (*entities.Product, error)
	UpdateProduct(ctx context.Context, product *entities.Product) error
	DeleteProduct(ctx context.Context, id string) error
	ListProducts(ctx context.Context, offset, limit int) ([]entities.Product, error)
}
