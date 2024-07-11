package envvars

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(dir string) {
	// Aquí se carga un archivo .env desde el directorio de trabajo actual.
	err := godotenv.Load(dir)
	if err != nil {
		log.Fatalf("Unable to load .env file")
	}
}
