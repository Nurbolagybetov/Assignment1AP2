package backend

import (
	"inventory-service/config"
	"inventory-service/internal/usecase"
)

type BackendServer interface {
	Run(port string) error
}

func NewBackendServer(cfg *config.Config, usecase usecase.ProductUsecase) BackendServer {
	return NewHTTPServer(cfg, usecase)
}
