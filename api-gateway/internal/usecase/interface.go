package usecase

import "context"

type GatewayUsecase interface {
	ForwardToInventory(ctx context.Context, endpoint string, method string, body []byte) ([]byte, error)
	ForwardToOrder(ctx context.Context, endpoint string, method string, body []byte) ([]byte, error)
}
