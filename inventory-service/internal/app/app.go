package app

import (
	"inventory-service/config"
	"inventory-service/internal/adapter/backend"
	"inventory-service/internal/mongo"
	"inventory-service/internal/usecase"
	"log"
)

type App struct {
	Config      *config.Config
	MongoClient *mongo.MongoClient
	Usecase     usecase.ProductUsecase
	Server      backend.BackendServer
}

func NewApp() *App {
	cfg := config.LoadConfig()
	mongoOpts := mongo.NewOptions(cfg.MongoURI, "inventory_db")
	mongoClient, err := mongo.NewMongoClient(mongoOpts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	repo := mongo.NewProductRepository(mongoClient)
	usecase := usecase.NewProductUsecase(repo)
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
	log.Printf("Inventory Service running on port %s", a.Config.Port)
	return a.Server.Run(a.Config.Port)
}
