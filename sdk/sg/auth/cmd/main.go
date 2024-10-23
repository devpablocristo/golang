package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"

	sdkaws "github.com/devpablocristo/golang/sdk/pkg/aws/localstack"
	sdkcnfldr "github.com/devpablocristo/golang/sdk/pkg/config/configLoader"
)

func init() {
	if err := sdkcnfldr.LoadConfig("config/.env", "config/.env.local"); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func awsInit() error {
	// AWS
	stack, err := sdkaws.Bootstrap()
	if err != nil {
		return err
	}

	// Usar la configuración de AWS
	awsCfg := stack.GetCfg()

	// Por ejemplo, crear un cliente para S3
	//s3Client
	_ = s3.NewFromConfig(awsCfg)

	fmt.Println("AWS stack initialized successfully with LocalStack")

	return nil
}

func main() {
	app, err := NewApplication()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	if err := app.Handler.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// func main() {
// 	// AWS
// 	stack, err := sdkaws.Bootstrap()
// 	if err != nil {
// 		log.Fatalf("Failed to bootstrap AWS stack: %v", err)
// 	}

// 	// Usar la configuración de AWS
// 	awsCfg := stack.GetCfg()

// 	// Por ejemplo, crear un cliente para S3
// 	//s3Client
// 	_ = s3.NewFromConfig(awsCfg)

// 	fmt.Println("AWS stack initialized successfully with LocalStack")
// 	// Fin AWS

// 	jwtService, err := authconn.NewJwtService()
// 	if err != nil {
// 		log.Fatalf("JWT Service error: %v", err)
// 	}

// 	httpClient, err := authconn.NewHttpClient()
// 	if err != nil {
// 		log.Fatalf("Http Client error: %v", err)
// 	}

// 	sessionManager, err := authconn.NewGorillaSessionManager()
// 	if err != nil {
// 		log.Fatalf("Http Client error: %v", err)
// 	}

// 	repository, err := authconn.NewPostgreSQL()
// 	if err != nil {
// 		log.Fatalf("PostgreSQL error: %v", err)
// 	}

// 	authUsecases := auth.NewUseCases(jwtService, repository, httpClient, sessionManager)

// 	authHandler, err := authgtw.NewGinHandler(authUsecases)
// 	if err != nil {
// 		log.Fatalf("Auth Handler error: %v", err)
// 	}

// 	err = authHandler.Start()
// 	if err != nil {
// 		log.Fatalf("Gin Server error at start: %v", err)
// 	}

// }
