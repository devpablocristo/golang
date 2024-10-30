package main

import (
	"fmt"
	"log"
	"time"

	sdkcnfldr "github.com/devpablocristo/golang/sdk/pkg/config/config-loader"
	"github.com/spf13/viper"

	mailconn "github.com/devpablocristo/golang/sdk/sg/mailing/internal/adapters/connectors"
	mailgtw "github.com/devpablocristo/golang/sdk/sg/mailing/internal/adapters/gateways"
	mailing "github.com/devpablocristo/golang/sdk/sg/mailing/internal/core"
)

func init() {
	if err := sdkcnfldr.LoadConfig("config/.env", "config/.env.local"); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}

	fmt.Println("SMTP Server:", viper.GetString("SMTP_SERVER"))
	fmt.Println("SMTP Port:", viper.GetString("SMTP_PORT"))
}

func main() {

	smtpService, err := mailconn.NewSmtpService()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	email := "pablo.cristo@teamcubation.com"
	expiration := 20 * time.Minute // El token de verificación expirará en 10 minutos

	err = smtpService.SendVerificationEmail(email, expiration)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	mailingUseCases := mailing.NewUseCases(smtpService)

	userHandler, err := mailgtw.NewGinHandler(mailingUseCases)
	if err != nil {
		log.Fatalf("Failed to initialize handler: %v", err)
	}

	fmt.Printf("Verification email sent to %s\n", email)

	err = userHandler.Start()
	if err != nil {
		log.Fatalf("Gin Server error at start: %v", err)
	}

}
