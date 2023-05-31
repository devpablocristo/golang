package main

import (
	"fmt"
	"os"

	envvars "testingdontenv/1/2/3/4"
)

// La ruta al archivo .env que proporcionamos a godotenv.Load() es relativa al directorio de trabajo actual
// del proceso, no necesariamente al archivo de código fuente donde se llama a godotenv.Load().
// Por tanto, el path correcto al archivo .env dependerá de desde dónde se esté ejecutando la aplicación:
// Si se lanza el programa desde la raíz del proyecto (es decir, se ejecuta 'go run cmd/app/main.go'
// estando en el directorio raíz), entonces la ruta al archivo .env debería ser simplemente ".env".
// Por otro lado, si se cambia primero al directorio cmd/app (mediante 'cd cmd/app') antes de ejecutar
// el programa (es decir, se ejecuta 'go run main.go' estando en el directorio cmd/app), entonces la ruta
// al archivo .env debería ser "../../.env", ya que ahora el directorio de trabajo es cmd/app y necesitamos
// retroceder dos niveles para llegar a la raíz del proyecto donde se encuentra el archivo .env.

func main() {
	envvars.LoadEnv("../../.env") // Carga las variables de entorno

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	fmt.Printf("DB_USER: %s, DB_PASSWORD: %s\n", dbUser, dbPassword)
}
