package main

import (
	"fmt"
	"os"

	envvars "testingdontenv/1/2/3/4"
)

// Depende desde donde estas parado al lanzar la correr la app
// se debera configurar un path o otro:
// si se corre desde la raiz del proyecto (con go run cmd/app/main.go )
// entonces el path debe ser simplmente ".env"
// si primero se cambia al directorio cmd/app (cd cmd/app),
// luego el path correcto sera "../../.env"
func main() {
	envvars.LoadEnv("../../.env") // Carga las variables de entorno

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	fmt.Printf("DB_USER: %s, DB_PASSWORD: %s\n", dbUser, dbPassword)
}
