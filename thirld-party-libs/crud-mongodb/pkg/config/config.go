package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// MongoDB configuration
	MongoUser     string
	MongoPassword string
	MongoHost     string
	MongoProtocol string
	DbName        string
	CollName      string
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

	MongoUser = os.Getenv("MONGODB_USER")
	MongoPassword = os.Getenv("MONGODB_PASSWORD")
	MongoHost = os.Getenv("MONGODB_HOST")
	MongoProtocol = os.Getenv("MONGODB_PROTOCOL")
	DbName = os.Getenv("MONGODB_DB_NAME")
	CollName = os.Getenv("MONGODB_COLLECTION_NAME")
}
