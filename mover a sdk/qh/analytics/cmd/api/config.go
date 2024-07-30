package api

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	rep "github.com/devpablocristo/qh/analytics/internal/core/report"
	stg "github.com/devpablocristo/qh/analytics/internal/platform/stage"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
}

type Dependencies struct {
	Repository rep.RepositoryPort
	DBConfig   DBConfig
	RouterPort string
}

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func getDBConfig(env stg.Stage) (DBConfig, error) {
	dbConfig := DBConfig{
		Host:     os.Getenv(fmt.Sprintf("%s_DB_HOST", env)),
		User:     os.Getenv(fmt.Sprintf("%s_DB_USERNAME", env)),
		Password: os.Getenv(fmt.Sprintf("%s_DB_USER_PASSWORD", env)),
	}

	if dbConfig.Host == "" || dbConfig.User == "" || dbConfig.Password == "" {
		return DBConfig{}, errors.New("incomplete database configuration")
	}

	return dbConfig, nil
}

func Config() (*Dependencies, error) {
	if err := loadEnv(); err != nil {
		return nil, err
	}

	appEnv := stg.GetFromString(os.Getenv("STAGE"))
	dbConfig, err := getDBConfig(appEnv)
	if err != nil {
		return nil, err
	}

	routerPort := os.Getenv("ROUTER_PORT")
	if routerPort == "" {
		return nil, errors.New("empty router port")
	}

	return &Dependencies{
		Repository: rep.NewRepository(),
		DBConfig:   dbConfig,
		RouterPort: routerPort,
	}, nil
}
