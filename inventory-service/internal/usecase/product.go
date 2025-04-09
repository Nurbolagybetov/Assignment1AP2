package usecase

import (
	"context"
	"inventory-service/internal/entities"
)

type productUsecase struct {
	repo ProductRepository
}

func NewProductUsecase(repo ProductRepository) ProductUsecase {
	return &productUsecase{repo: repo}
}

func (u *productUsecase) CreateProduct(ctx context.Context, product *entities.Product) error {
	return u.repo.Create(ctx, product)
}

func (u *productUsecase) GetProduct(ctx context.Context, id string) (*entities.Product, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *productUsecase) UpdateProduct(ctx context.Context, product *entities.Product) error {
	return u.repo.Update(ctx, product)
}

func (u *productUsecase) DeleteProduct(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}

func (u *productUsecase) ListProducts(ctx context.Context, offset, limit int) ([]entities.Product, error) {
	return u.repo.List(ctx, offset, limit)
}
