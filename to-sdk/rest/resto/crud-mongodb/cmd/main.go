// Este código es un ejemplo básico de cómo conectarse a MongoDB,
// consultar una colección y mostrar los resultados en la consola utilizando el controlador MongoDB para Go
// (también conocido como mongo-go-driver).
package main

import (
	"context"
	"log"
	"time"

	mongodbrepo "crudmongodb/internal/adapters/driven/repositories/mongodb"
	service "crudmongodb/internal/service"
)

func main() {
	// Crea un contexto con un tiempo de espera de 50 segundos.
	// El contexto se utiliza para definir cuánto tiempo debe esperar el programa
	// antes de que una operación de MongoDB se considere fallida. defer cancel()
	// garantiza que el contexto se cancele cuando la función main termine.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	client, err := mongodbrepo.ConnectMongoDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	repo := mongodbrepo.NewMongoService(client, ctx)
	serv := service.NewService(repo, ctx)

	// newDocument := domain.Listing{
	// 	Name: "Nuevo alojamiento",
	// }

	//serv.Create(newDocument)
	serv.ReadAll()
	// serv.Update()
	// serv.Delete()

}
