package usecase

import (
	"context"
	"order-service/internal/entities"
)

type OrderRepository interface {
	Create(ctx context.Context, order *entities.Order) error
	GetByID(ctx context.Context, id string) (*entities.Order, error)
	Update(ctx context.Context, order *entities.Order) error
	List(ctx context.Context, userID string) ([]entities.Order, error)
}

type OrderUsecase interface {
	CreateOrder(ctx context.Context, order *entities.Order) error
	GetOrder(ctx context.Context, id string) (*entities.Order, error)
	UpdateOrder(ctx context.Context, order *entities.Order) error
	ListOrders(ctx context.Context, userID string) ([]entities.Order, error)
}
