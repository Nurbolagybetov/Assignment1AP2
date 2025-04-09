package backend

import (
	"order-service/config"
	"order-service/internal/usecase"
)

type BackendServer interface {
	Run(port string) error
}

func NewBackendServer(cfg *config.Config, usecase usecase.OrderUsecase) BackendServer {
	return NewHTTPServer(cfg, usecase)
}
