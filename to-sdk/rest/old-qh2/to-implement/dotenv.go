package application

import (
	"fmt"
	"log"

	godotenv "github.com/joho/godotenv"
)

var (
	url_help     = "URL_HELP"
	url_glossary = "URL_GLOSSARY"
)

func loadENVConfigs() (map[string]string, error) {
	env, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return env, nil
}

func setupDotEnv() {

	envs, err := loadENVConfigs()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	loadENVConfigs()

	name := envs[url_glossary]
	editor := envs[url_help]

	fmt.Printf("%s uses %s\n", name, editor)
}
