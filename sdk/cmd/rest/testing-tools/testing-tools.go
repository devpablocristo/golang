package tttools

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// TODO: normalizar y mover a pkg
func LoadTestData(filename string) ([]byte, error) {
	// Obtener el directorio base del archivo actual
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to obtain caller information")
	}

	basePath := filepath.Dir(b)
	filePath := filepath.Join(basePath, "data", filename)

	// Leer los datos del archivo
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	return fileData, nil
}
