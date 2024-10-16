package shared

import (
	"fmt"
	"os"
	"path/filepath"
)

func LoadTestData(dir, filename string) ([]byte, error) {
	// Crear la ruta completa usando el directorio y el nombre del archivo
	filePath := filepath.Join(dir, filename)

	// Leer los datos del archivo
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	return fileData, nil
}
