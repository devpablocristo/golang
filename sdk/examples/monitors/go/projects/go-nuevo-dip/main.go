package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"

	sdkgodotenv "github.com/devpablocristo/golang/sdk/pkg/configurators/godotenv"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgogit "github.com/devpablocristo/golang/sdk/pkg/repo-tools/go-git/v5"
	sdkff "github.com/devpablocristo/golang/sdk/pkg/tools/files-finder"
)

type LayersConfig struct {
	Layers struct {
		Domain         []string `mapstructure:"domain"`
		Application    []string `mapstructure:"application"`
		Infrastructure []string `mapstructure:"infrastructure"`
	} `mapstructure:"layers"`
}

func main() {
	envFile, err := sdkff.FilesFinder("config/.env")
	if err != nil {
		fmt.Printf("Error finding files: %v\n", err)
	}

	ymlFile, err := sdkff.FilesFinder("services/monitors/go/projects/go-nuevo-dip/layers.yml")
	if err != nil {
		fmt.Printf("Error finding files: %v\n", err)
	}

	err = sdkgodotenv.LoadConfig(envFile...)
	if err != nil {
		log.Fatalf("Error loading .env files: %v", err)
	}

	err = sdkviper.LoadConfig(ymlFile...)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	appName := viper.GetString("APP_NAME")
	debug := viper.GetBool("DEBUG")
	fmt.Printf("App Name: %s, Debug Mode: %t\n", appName, debug)

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <repo_path>")
		return
	}

	repoPath := os.Args[1]

	//repo, err := sdkgogit.Bootstrap("https://github.com/devpablocristo/meli", "/home/pablo/meli", "main")
	repo, err := sdkgogit.Bootstrap(repoPath, "/home/pablo/meli", "main")
	if err != nil {
		log.Fatalf("Error initializing Git client: %v", err)
	}

	err = repo.PullLatest()
	if err != nil {
		log.Fatalf("Error pulling latest changes: %v", err)
	}

	layersConfig, err := loadLayerConfig()
	if err != nil {
		log.Fatalf("Error loading layer configuration: %v", err)
	}

	fmt.Println("Layers Configuration:")
	fmt.Printf("Domain Layers: %v\n", layersConfig.Layers.Domain)
	fmt.Printf("Application Layers: %v\n", layersConfig.Layers.Application)
	fmt.Printf("Infrastructure Layers: %v\n", layersConfig.Layers.Infrastructure)

	files, err := repo.GetFiles(nil, ".go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)

	fmt.Println(repo.GetFileAuthor(files[0]))
	fmt.Println(repo.GetCommitID(files[0]))

}
