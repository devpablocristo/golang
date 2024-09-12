package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Definición de la estructura Metric
type Metric struct {
	MetricID  string     `json:"metric_id"`  // Identificador de la métrica
	GitAuthor string     `json:"git_author"` // Autor del commit
	Score     int        `json:"score"`      // Puntaje de cumplimiento
	Evidence  []Evidence `json:"evidence"`   // Evidencias recolectadas
}

// Definición de la estructura Evidence
type Evidence struct {
	CommitID string `json:"commit_id"` // ID del commit asociado
	File     string `json:"file"`      // Archivo donde se encuentra la evidencia
	Line     int    `json:"line"`      // Línea del archivo donde está la violación o evidencia
}

// Definimos una variable global de tipo Metric con un puntaje inicial
var globalMetric = Metric{
	MetricID: "dependency_inversion",
	Score:    3, // Puntaje inicial
	Evidence: []Evidence{},
}

var baseDir string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <repo_path> [file1] [file2] ...")
		return
	}

	repoPath := os.Args[1]
	filesToAnalyze := os.Args[2:]

	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Printf("Error opening repository: %v\n", err)
		return
	}

	head, err := repo.Head()
	if err != nil {
		fmt.Printf("Error getting HEAD: %v\n", err)
		return
	}

	commit, err := repo.CommitObject(head.Hash())
	if err != nil {
		fmt.Printf("Error getting commit: %v\n", err)
		return
	}

	tree, err := commit.Tree()
	if err != nil {
		fmt.Printf("Error getting tree: %v\n", err)
		return
	}

	baseDir = repoPath

	var files []string
	if len(filesToAnalyze) == 0 {
		err = tree.Files().ForEach(func(f *object.File) error {
			if filepath.Ext(f.Name) == ".go" {
				files = append(files, f.Name)
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error iterating files: %v\n", err)
			os.Exit(1)
		}
	} else {
		for _, file := range filesToAnalyze {
			if filepath.Ext(file) == ".go" {
				files = append(files, file)
			}
		}
	}

	if len(files) == 0 {
		fmt.Println("No files to analyze.")
		return
	}

	metrics := analyzeCode(repo, tree, files)

	jsonOutput, err := json.MarshalIndent(metrics, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonOutput))
}

func analyzeCode(repo *git.Repository, tree *object.Tree, filesToAnalyze []string) []Metric {
	var metrics []Metric

	author, err := getFileAuthor(repo, filesToAnalyze[0])
	if err != nil {
		fmt.Printf("Error getting file author: %v\n", err)
		os.Exit(1)
	}

	dipMetric, err := analyzeDependencyInversion(repo, tree, filesToAnalyze)
	if err != nil {
		log.Fatal(err)
	}

	dipMetric.GitAuthor = author
	metrics = append(metrics, dipMetric)

	return metrics
}

func analyzeDependencyInversion(repo *git.Repository, tree *object.Tree, filesToAnalyze []string) (Metric, error) {
	// Aquí usamos la métrica global
	metric := globalMetric

	interfaces := make(map[string]bool)
	concreteDeps := false

	for _, filePath := range filesToAnalyze {
		content, err := os.ReadFile(filepath.Join(baseDir, filePath))
		if err != nil {
			return metric, fmt.Errorf("error reading file: %v", err)
		}

		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, filePath, content, parser.AllErrors)
		if err != nil {
			return metric, fmt.Errorf("error parsing file: %v", err)
		}

		ast.Inspect(node, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.TypeSpec:
				// Revisamos si hay interfaces definidas
				if _, ok := x.Type.(*ast.InterfaceType); ok {
					interfaces[x.Name.Name] = true
				}
			case *ast.CallExpr:
				// Verificamos si se llama a una función que NO es una interfaz
				if ident, ok := x.Fun.(*ast.Ident); ok {
					if !interfaces[ident.Name] {
						concreteDeps = true // Se detecta una dependencia concreta
						metric.Evidence = append(metric.Evidence, Evidence{
							File: filePath,
							Line: fset.Position(x.Pos()).Line,
						})
					}
				}
			}
			return true
		})
	}

	// Ajustamos el puntaje dependiendo de las condiciones
	if len(interfaces) == 0 && concreteDeps {
		metric.Score = 1 // No hay interfaces y se encontraron dependencias concretas (Violación completa de DIP)
	} else if len(interfaces) > 0 && concreteDeps {
		metric.Score = 2 // Hay interfaces, pero también hay dependencias concretas (Violación parcial de DIP)
	} else {
		metric.Score = 3 // Todo está bien, no hay dependencias concretas
	}

	return metric, nil
}

func getFileAuthor(repo *git.Repository, file string) (string, error) {
	commits, err := repo.Log(&git.LogOptions{FileName: &file})
	if err != nil {
		return "", err
	}

	commit, err := commits.Next()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s <%s>", commit.Author.Name, commit.Author.Email), nil
}
