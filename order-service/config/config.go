package config

type Config struct {
	Port     string
	MongoURI string
}

func LoadConfig() *Config {
	return &Config{
		Port:     ":8082",
		MongoURI: "mongodb://localhost:27017",
	}
}
