package main

import (
	"fmt"
	"log"

	sdkawslocal "github.com/devpablocristo/golang/sdk/pkg/aws/localstack"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
)

func init() {
	if err := sdkviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
	// Bootstrap del servicio
	service, err := sdkawslocal.Bootstrap()
	if err != nil {
		log.Fatalf("Failed to bootstrap AWS service: %v", err)
	}

	// Usar la configuración de AWS
	_ = service.GetAWSCfg()

	// Por ejemplo, crear un cliente para S3
	// s3Client := s3.NewFromConfig(awsCfg)

	fmt.Println("AWS service initialized successfully with LocalStack")
}
