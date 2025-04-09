package config

type Config struct {
	Port         string
	InventoryURL string
	OrderURL     string
	MongoURI     string
}

func LoadConfig() *Config {
	return &Config{
		Port:         ":8080",
		InventoryURL: "http://localhost:8081",
		OrderURL:     "http://localhost:8082",
		MongoURI:     "mongodb://localhost:27017",
	}
}
