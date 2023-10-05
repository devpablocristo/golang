package config

import "os"

type Config struct {
	GRPCAddress   string
	MongoURI      string
	MongoDatabase string
}

func LoadConfig() Config {
	return Config{
		GRPCAddress:   getEnv("GRPC_ADDRESS", ":50051"),
		MongoURI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDatabase: getEnv("MONGO_DATABASE", "chat"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
