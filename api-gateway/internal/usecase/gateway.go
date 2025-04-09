package usecase

import (
	"api-gateway/config"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

type gatewayUsecase struct {
	cfg *config.Config
}

func NewGatewayUsecase(cfg *config.Config) GatewayUsecase {
	return &gatewayUsecase{cfg: cfg}
}

func (u *gatewayUsecase) ForwardToInventory(ctx context.Context, endpoint, method string, body []byte) ([]byte, error) {
	url := fmt.Sprintf("%s%s", u.cfg.InventoryURL, endpoint)
	return u.forwardRequest(ctx, url, method, body)
}

func (u *gatewayUsecase) ForwardToOrder(ctx context.Context, endpoint, method string, body []byte) ([]byte, error) {
	url := fmt.Sprintf("%s%s", u.cfg.OrderURL, endpoint)
	return u.forwardRequest(ctx, url, method, body)
}

func (u *gatewayUsecase) forwardRequest(ctx context.Context, url, method string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
