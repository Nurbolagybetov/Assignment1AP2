package config

type Config struct {
	Port     string
	MongoURI string
}

func LoadConfig() *Config {
	return &Config{
		Port:     ":8081",
		MongoURI: "mongodb://localhost:27017",
	}
}
