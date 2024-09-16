package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Estructura para almacenar información de interfaces encontradas
type Interface struct {
	Name string
	Path string
}

// Función para leer y analizar archivos .go para encontrar interfaces y variables
func analyzeGoFile(filePath string, interfaces *[]Interface) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Regex para detectar la declaración de interfaces y variables
	interfaceRegex := regexp.MustCompile(`type\s+(\w+)\s+interface`)
	varDeclRegex := regexp.MustCompile(`var\s+(\w+)\s+(\w+)`)

	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()

		// Buscar interfaces declaradas
		if interfaceMatch := interfaceRegex.FindStringSubmatch(line); interfaceMatch != nil {
			interf := Interface{
				Name: interfaceMatch[1],
				Path: filePath,
			}
			*interfaces = append(*interfaces, interf)
			fmt.Printf("Interface encontrada: %s en archivo %s, línea %d\n", interf.Name, filePath, lineNumber)
		}

		// Buscar variables declaradas
		if varMatch := varDeclRegex.FindStringSubmatch(line); varMatch != nil {
			fmt.Printf("Variable encontrada: %s, Tipo: %s en archivo %s, línea %d\n", varMatch[1], varMatch[2], filePath, lineNumber)
		}

		lineNumber++
	}

	return scanner.Err()
}

// Función para escanear el repositorio y encontrar archivos .go con interfaces y variables
func scanRepoForInterfacesAndVars(repoPath string, interfaces *[]Interface) error {
	err := filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Procesar solo archivos .go
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			err := analyzeGoFile(path, interfaces)
			if err != nil {
				fmt.Printf("Error analizando archivo %s: %v\n", path, err)
			}
		}
		return nil
	})

	return err
}

// Función para encontrar el uso de interfaces en cada archivo
func findInterfaceUsage(filePath string, interfaces []Interface) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	var foundUsages []string

	for scanner.Scan() {
		line := scanner.Text()

		// Verificar si alguna interfaz está siendo utilizada
		for _, interf := range interfaces {
			if strings.Contains(line, interf.Name) {
				foundUsages = append(foundUsages, fmt.Sprintf("Uso de interfaz %s en línea %d", interf.Name, lineNumber))
			}
		}

		lineNumber++
	}

	// Solo imprimir si se encontraron usos de interfaces en este archivo
	if len(foundUsages) > 0 {
		fmt.Printf("\nArchivo: %s\n", filePath)
		for _, usage := range foundUsages {
			fmt.Println(usage)
		}
		fmt.Println("-----------------------------")
	}

	return scanner.Err()
}

// Función para buscar el uso de interfaces en todo el repositorio
func scanRepoForInterfaceUsage(repoPath string, interfaces []Interface) error {
	err := filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Procesar solo archivos .go
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			err := findInterfaceUsage(path, interfaces)
			if err != nil {
				fmt.Printf("Error buscando uso de interfaces en archivo %s: %v\n", path, err)
			}
		}
		return nil
	})

	return err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <ruta_del_repositorio>")
		return
	}

	repoPath := os.Args[1]

	// Almacenar interfaces encontradas
	var interfaces []Interface

	// Escanear el repositorio para encontrar interfaces y variables
	fmt.Println("Buscando interfaces y variables en el repositorio...")
	err := scanRepoForInterfacesAndVars(repoPath, &interfaces)
	if err != nil {
		fmt.Printf("Error al escanear el repositorio: %v\n", err)
		return
	}

	// Buscar el uso de las interfaces en el repositorio
	fmt.Println("Buscando uso de interfaces en el repositorio...")
	err = scanRepoForInterfaceUsage(repoPath, interfaces)
	if err != nil {
		fmt.Printf("Error al escanear el uso de interfaces en el repositorio: %v\n", err)
	}
}
