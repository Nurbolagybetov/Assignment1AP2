package usecase

import (
	"context"
	"order-service/internal/entities"
)

type orderUsecase struct {
	repo OrderRepository
}

func NewOrderUsecase(repo OrderRepository) OrderUsecase {
	return &orderUsecase{repo: repo}
}

func (u *orderUsecase) CreateOrder(ctx context.Context, order *entities.Order) error {
	return u.repo.Create(ctx, order)
}

func (u *orderUsecase) GetOrder(ctx context.Context, id string) (*entities.Order, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *orderUsecase) UpdateOrder(ctx context.Context, order *entities.Order) error {
	return u.repo.Update(ctx, order)
}

func (u *orderUsecase) ListOrders(ctx context.Context, userID string) ([]entities.Order, error) {
	return u.repo.List(ctx, userID)
}
