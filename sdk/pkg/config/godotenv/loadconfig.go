package sdkgodotenv

import (
	"fmt"

	"github.com/joho/godotenv"

	sdktools "github.com/devpablocristo/golang/sdk/pkg/tools"
)

// LoadConfig carga múltiples archivos desde las rutas proporcionadas.
// Cada string en configPaths debe representar el nombre del archivo o ruta relativa.
// Retorna un error en caso de que no se pueda cargar algún archivo.
func LoadConfig(configPaths ...string) error {
	// Si no se proporcionan rutas, se debe intentar buscar un archivo por defecto.
	if len(configPaths) == 0 {
		configPaths = []string{".env"}
	}

	// Buscar los archivos en el proyecto utilizando FilesFinder
	foundFiles, err := sdktools.FilesFinder(configPaths...)
	if err != nil {
		return fmt.Errorf("FilesFinder failed to find files: %w", err)
	}

	// Cargar cada archivo desde las rutas encontradas
	for _, filePath := range foundFiles {
		fmt.Printf("SDK Envgodot: Attempting to load file from: %s\n", filePath) // Debugging
		// Aquí no necesitas procesar de nuevo la ruta, simplemente la pasas a godotenv.Load
		if err := godotenv.Load(filePath); err != nil {
			return fmt.Errorf("failed to load file from path %s: %w", filePath, err)
		}
	}

	fmt.Println("SDK Envgodot: Files loaded successfully.")
	return nil
}
