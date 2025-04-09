package app

import (
	"log"
	"order-service/config"
	"order-service/internal/adapter/backend"
	"order-service/internal/mongo"
	"order-service/internal/usecase"
)

type App struct {
	Config      *config.Config
	MongoClient *mongo.MongoClient
	Usecase     usecase.OrderUsecase
	Server      backend.BackendServer
}

func NewApp() *App {
	cfg := config.LoadConfig()
	mongoOpts := mongo.NewOptions(cfg.MongoURI, "order_db")
	mongoClient, err := mongo.NewMongoClient(mongoOpts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	repo := mongo.NewOrderRepository(mongoClient)
	usecase := usecase.NewOrderUsecase(repo)
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
	log.Printf("Order Service running on port %s", a.Config.Port)
	return a.Server.Run(a.Config.Port)
}
