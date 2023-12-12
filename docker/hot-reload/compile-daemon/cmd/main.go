package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	handler "compile-daemon/infra/handler"
)

func main() {
	// Configuración de AWS DynamoDB
	config := &aws.Config{
		Endpoint: aws.String("http://dynamodb-local:8765"), // Nombre del servicio y puerto de Docker Compose
		Region:   aws.String("local"),                      // Usa 'local' para DynamoDB local
		Credentials: credentials.NewStaticCredentials(
			"fakeAccessKey",
			"fakeSecretKey",
			"",
		),
	}

	// Crear una nueva sesión
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	// Crear un cliente DynamoDB
	svc := dynamodb.New(sess)

	// Ejemplo: Listar tablas
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Panicf("error: %v", err)
	}

	fmt.Println("Tablas en DynamoDB:")
	for _, table := range result.TableNames {
		fmt.Println(*table)
	}

	http.HandleFunc("/", handler.HomePage)
	http.HandleFunc("/users", handler.UserPage)
	// puerto del contenedor
	log.Fatal(http.ListenAndServe(":8888", nil))
}
