package backend

import (
	"api-gateway/config"
	"api-gateway/internal/usecase"
)

type BackendServer interface {
	Run(port string) error
}

func NewBackendServer(cfg *config.Config, usecase usecase.GatewayUsecase) BackendServer {
	return NewHTTPServer(cfg, usecase)
}
