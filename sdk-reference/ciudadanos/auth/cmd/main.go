package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"

	sdkawslocal "github.com/devpablocristo/golang/sdk/pkg/aws/localstack"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"

	authconn "github.com/devpablocristo/golang/sdk/ciudadanos/auth/internal/adapters/connectors"
	authgtw "github.com/devpablocristo/golang/sdk/ciudadanos/auth/internal/adapters/gateways"
	auth "github.com/devpablocristo/golang/sdk/ciudadanos/auth/internal/core"
)

func init() {
	if err := sdkviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {

	////////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////////

	jwtService, err := authconn.NewJwtService()
	if err != nil {
		log.Fatalf("JWT Service error: %v", err)
	}

	authUsecases := auth.NewUseCases(jwtService)

	authHandler, err := authgtw.NewGinHandler(authUsecases)
	if err != nil {
		log.Fatalf("Auth Handler error: %v", err)
	}

	err = authHandler.Start()
	if err != nil {
		log.Fatalf("Gin Server error at start: %v", err)
	}

	// AWS
	stack, err := sdkawslocal.Bootstrap()
	if err != nil {
		log.Fatalf("Failed to bootstrap AWS stack: %v", err)
	}

	// Usar la configuración de AWS
	awsCfg := stack.GetCfg()

	// Por ejemplo, crear un cliente para S3
	//s3Client
	_ = s3.NewFromConfig(awsCfg)

	fmt.Println("AWS stack initialized successfully with LocalStack")

}
