package api

import (
	"errors"
	"os"

	"github.com/joho/godotenv"

	ltp "github.com/devpablocristo/dive-challenge/internal/core/ltp"
	stg "github.com/devpablocristo/dive-challenge/internal/platform/stage"
)

type Dependencies struct {
	Repository ltp.RepositoryPort
	ApiClient  ltp.APIClientPort
	DBHost     string
	DBUser     string
	DBPassword string
	RouterPort string
}

func Config() (*Dependencies, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("error loading .env file: " + err.Error())
	}

	appEnv := stg.GetFromString(os.Getenv("STAGE"))

	var dbHost, dbUser, dbPassword string
	switch appEnv {
	case stg.Production:
		dbHost = os.Getenv("PROD_DB_HOST")
		dbUser = os.Getenv("PROD_DB_USER")
		dbPassword = os.Getenv("PROD_DB_PASSWORD")
	case stg.Beta:
		dbHost = os.Getenv("BETA_DB_HOST")
		dbUser = os.Getenv("BETA_DB_USER")
		dbPassword = os.Getenv("BETA_DB_PASSWORD")
	case stg.Development:
		dbHost = os.Getenv("DEV_DB_HOST")
		dbUser = os.Getenv("DEV_DB_USERNAME")
		dbPassword = os.Getenv("DEV_DB_USER_PASSWORD")
	}

	if dbHost == "" || dbUser == "" || dbPassword == "" {
		return nil, errors.New("database configuration is incomplete")
	}

	routerPort := os.Getenv("ROUTER_PORT")
	if routerPort == "" {
		return nil, errors.New("empty router port")
	}

	//TODO: change url
	url := "https://api.kraken.com"

	return &Dependencies{
		Repository: ltp.NewRepository(),
		ApiClient:  ltp.NewExternalAPI(url),
		DBHost:     dbHost,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		RouterPort: routerPort,
	}, nil
}
