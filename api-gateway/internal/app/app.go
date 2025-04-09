package app

import (
	"api-gateway/config"
	"api-gateway/internal/adapter/backend"
	"api-gateway/internal/mongo"
	"api-gateway/internal/usecase"
	"log"
)

type App struct {
	Config      *config.Config
	MongoClient *mongo.MongoClient
	Usecase     usecase.GatewayUsecase
	Server      backend.BackendServer
}

func NewApp() *App {
	cfg := config.LoadConfig()
	mongoOpts := mongo.NewOptions(cfg.MongoURI, "api_gateway_db")
	mongoClient, err := mongo.NewMongoClient(mongoOpts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	usecase := usecase.NewGatewayUsecase(cfg)
	server := backend.NewBackendServer(cfg, usecase)

	return &App{
		Config:      cfg,
		MongoClient: mongoClient,
		Usecase:     usecase,
		Server:      server,
	}
}

func (a *App) Run() error {
	defer a.MongoClient.Disconnect()
	log.Printf("API Gateway running on port %s", a.Config.Port)
	return a.Server.Run(a.Config.Port)
}
