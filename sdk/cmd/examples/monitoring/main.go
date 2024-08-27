package main

import (
	"log"

	gtwmon "github.com/devpablocristo/golang/sdk/cmd/gateways/monitoring"
	coremon "github.com/devpablocristo/golang/sdk/internal/core/monitoring"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkmysql "github.com/devpablocristo/golang/sdk/pkg/databases/mysql/go-sql-driver"
	portsmysql "github.com/devpablocristo/golang/sdk/pkg/databases/mysql/go-sql-driver/ports"
	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
	portsgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
)

// CMND: docker compose -f ../../../config/docker-compose.dev.yml up --build
func init() {
	if err := sdkviper.LoadConfig("../../../"); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
	ginServer, mysqlRepository := lauchBootstraps()

	monRepository := coremon.NewMySqlRepository(mysqlRepository)
	userUsecases := coremon.NewUserUseCases(monRepository)
	userHandler := gtwmon.NewGinHandler(userUsecases, ginServer)
	userHandler.Start("v1")

}

func lauchBootstraps() (portsgin.Server, portsmysql.Repository) {
	mysqlRepository, err := sdkmysql.Bootstrap()
	if err != nil {
		log.Fatalf("MySQL Service error: %v", err)
	}

	ginServer, err := sdkgin.Bootstrap()
	if err != nil {
		log.Fatalf("Gin Service error: %v", err)
	}

	return ginServer, mysqlRepository
}
