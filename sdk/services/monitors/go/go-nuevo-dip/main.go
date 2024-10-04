package main

import (
	"fmt"

	sdkff "github.com/devpablocristo/golang/sdk/pkg/tools/files-finder"
)

// sdkgodotenv "github.com/devpablocristo/golang/sdk/pkg/configurators/godotenv"
// sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
// sdkgogit "github.com/devpablocristo/golang/sdk/pkg/repo-tools/go-git/v5"

// LayersConfig representa la configuración de capas en el archivo YAML.
type LayersConfig struct {
	Layers struct {
		Domain         []string `mapstructure:"domain"`
		Application    []string `mapstructure:"application"`
		Infrastructure []string `mapstructure:"infrastructure"`
	} `mapstructure:"layers"`
}

func main() {

	files, _ := sdkff.FilesFinder("config/.env", "services/monitors/go/go-nuevo-dip/layers.yml")

	fmt.Println(files)
	// // Ruta al archivo .env
	// envFilePath := "../../../../config"

	// // Ruta al archivo de configuración (por ejemplo, layers.yaml)
	// configFilePath := "./layers.yml" // Ajusta esto a la ruta correcta de tu archivo

	// // 1. Cargar variables de entorno desde archivos .env usando sdkgodotenv
	// err := sdkgodotenv.LoadConfig(envFilePath)
	// if err != nil {
	// 	log.Fatalf("Error loading .env files: %v", err)
	// }

	// // 2. Cargar configuraciones usando sdkviper
	// err = sdkviper.LoadConfig(configFilePath)
	// if err != nil {
	// 	log.Fatalf("Error loading configuration: %v", err)
	// }

	// // 3. Validar que las variables de entorno esenciales estén establecidas
	// requiredVars := []string{"APP_NAME", "DEBUG"}
	// err = sdkviper.ValidateEnvVars(requiredVars)
	// if err != nil {
	// 	log.Fatalf("Environment validation failed: %v", err)
	// }

	// // 4. Acceder a las variables de entorno y configuración
	// appName := viper.GetString("APP_NAME")
	// debug := viper.GetBool("DEBUG")
	// fmt.Printf("App Name: %s, Debug Mode: %t\n", appName, debug)

	// // 5. Obtener la configuración mapeada
	// var layersConfig LayersConfig
	// err = sdkviper.UnmarshalConfig(&layersConfig)
	// if err != nil {
	// 	log.Fatalf("Error unmarshaling configuration: %v", err)
	// }

	// // 6. Acceder a la configuración de layers
	// fmt.Println("Layers Configuration:")
	// fmt.Printf("Domain Layers: %v\n", layersConfig.Layers.Domain)
	// fmt.Printf("Application Layers: %v\n", layersConfig.Layers.Application)
	// fmt.Printf("Infrastructure Layers: %v\n", layersConfig.Layers.Infrastructure)

	// // 7. Inicializar el cliente Git usando sdkgogit
	// goGitClient, err := sdkgogit.Bootstrap()
	// if err != nil {
	// 	log.Fatalf("Error initializing Git client: %v", err)
	// }

	// // 8. Ejecutar un pull para obtener los últimos cambios del repositorio
	// err = goGitClient.PullLatest()
	// if err != nil {
	// 	log.Fatalf("Error pulling latest changes: %v", err)
	// }

	// fmt.Println("Repository updated successfully")
}
