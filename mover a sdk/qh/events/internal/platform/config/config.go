package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go-micro.dev/v4/logger"

	stg "github.com/devpablocristo/qh/events/internal/platform/stage"
)

// DBConfig contiene la configuración de la base de datos
type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
}

type ConsulConfig struct {
	ID      string
	Name    string
	Port    int
	Address string
	HTTP    string
	Service string
}

type RouterConfig struct {
	RouterPort string
}

// Dependencies contiene todas las dependencias configuradas
type Dependencies struct {
	DBConfig     DBConfig
	RouterConfig RouterConfig
	ConsulConfig ConsulConfig
}

// initConfig inicializa la configuración usando Viper
func initConfig() error {
	viper.SetConfigName(".env") // Nombre del archivo de configuración (sin extensión)
	viper.SetConfigType("env")  // Tipo de archivo de configuración
	viper.AddConfigPath(".")    // Ruta al directorio del archivo de configuración

	// Leer variables de entorno
	viper.AutomaticEnv()

	// Intentar leer el archivo de configuración
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	return nil
}

// initLogger inicializa el logger
func initLogger() error {
	err := logger.Init(
		logger.WithLevel(logger.InfoLevel), // Establecer el nivel de logging
		logger.WithOutput(os.Stdout),       // Establecer la salida a la consola
	)
	if err != nil {
		return err
	}
	return nil
}

// Load carga todas las dependencias configuradas
func Load() (*Dependencies, error) {
	if err := initLogger(); err != nil {
		return nil, fmt.Errorf("error initializing logger: %w", err)
	}

	if err := initConfig(); err != nil {
		return nil, fmt.Errorf("error initializing config: %w", err)
	}

	appEnv := stg.GetFromString(viper.GetString("STAGE"))
	if appEnv == stg.Unknown {
		return nil, errors.New("invalid or missing STAGE environment variable")
	}

	dbConfig := DBConfig{
		Host:     viper.GetString(fmt.Sprintf("%s_DB_HOST", appEnv)),
		User:     viper.GetString(fmt.Sprintf("%s_DB_USERNAME", appEnv)),
		Password: viper.GetString(fmt.Sprintf("%s_DB_USER_PASSWORD", appEnv)),
		DBName:   viper.GetString(fmt.Sprintf("%s_DB_DATABASE", appEnv)),
	}

	if dbConfig.Host == "" || dbConfig.User == "" || dbConfig.Password == "" || dbConfig.DBName == "" {
		return nil, errors.New("incomplete database configuration")
	}

	routerPort := viper.GetString("ROUTER_PORT")
	if routerPort == "" {
		return nil, errors.New("empty router port")
	}

	routerConfig := RouterConfig{
		RouterPort: routerPort,
	}

	consulConfig := ConsulConfig{
		ID:      viper.GetString("APP_NAME"),
		Name:    viper.GetString("APP_NAME"),
		Port:    viper.GetInt("ROUTER_PORT"),
		Address: viper.GetString("CONSUL_ADDRESS"),
		HTTP:    viper.GetString("CONSUL_HTTP"),
		Service: viper.GetString("CONSUL_SERVICE_NAME"),
	}

	if consulConfig.ID == "" || consulConfig.Name == "" || consulConfig.Port == 0 || consulConfig.Address == "" || consulConfig.HTTP == "" || consulConfig.Service == "" {
		return nil, errors.New("incomplete Consul configuration")
	}

	return &Dependencies{
		DBConfig:     dbConfig,
		RouterConfig: routerConfig,
		ConsulConfig: consulConfig,
	}, nil
}
