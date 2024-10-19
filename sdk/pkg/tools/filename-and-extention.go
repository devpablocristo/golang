package sdktools

import (
	"fmt"
	"path/filepath"
	"strings"
)

func FileNameAndExtension(filePath string) (string, string, error) {
	fileName := filepath.Base(filePath)

	// Verificar si el archivo comienza con un punto y no tiene una extensión real
	if strings.HasPrefix(fileName, ".") && strings.Count(fileName, ".") == 1 {
		// Caso donde el archivo es algo como ".env", ".gitignore", etc.
		return fileName, "", nil
	}

	// Obtener la extensión del archivo
	fileExtension := filepath.Ext(fileName)

	// Si no tiene extensión (archivo sin punto o solo con nombre), devolvemos error
	if fileExtension == "" {
		return "", "", fmt.Errorf("file %s has no extension", fileName)
	}

	// Eliminar el punto de la extensión
	fileExtension = strings.TrimPrefix(fileExtension, ".")

	// Obtener el nombre del archivo sin la extensión
	fileNameWithoutExt := strings.TrimSuffix(fileName, "."+fileExtension)

	// Si el archivo no tiene un nombre válido, devolver error
	if fileNameWithoutExt == "" {
		return "", "", fmt.Errorf("invalid file name for file %s", fileName)
	}

	return fileNameWithoutExt, fileExtension, nil
}
