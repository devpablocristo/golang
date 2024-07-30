package main

import (
	"context"
	"log"

	application "github.com/luka385/tinder-pets/application/user_service"
	http2 "github.com/luka385/tinder-pets/infrastructure/http"
	"github.com/luka385/tinder-pets/infrastructure/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ClientOption := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), ClientOption)
	if err != nil {
		log.Fatal("Error al conectar MongoDB:", err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("test")
	userRepository := mongodb.NewUserRepository(db)
	userService := application.NewUserService(userRepository)

	r := http2.SetupsServer(userService)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("error al conectar el servidor", err)
	}
}
