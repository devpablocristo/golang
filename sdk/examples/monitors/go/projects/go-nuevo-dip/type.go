package main

// ---------------------------
// Estructuras principales y análisis de dependencias
// ---------------------------

// EntityInfo representa información sobre una entidad dentro de un archivo (struct, variable, etc.)
type EntityInfo struct {
	Name        string
	Type        string
	Position    int
	Category    string
	Kind        string
	IsInterface bool
	Layer       string
}

// FileImport representa un archivo Go y las entidades que contiene
type FileImport struct {
	Name     string
	Path     string
	Layer    string
	Entities []EntityInfo
}

// DependencyAnalyzer es la estructura principal para analizar dependencias
type DependencyAnalyzer struct {
	PackagesInfo         []FileImport
	DependencyViolations []string
}

// ---------------------------
// Estructuras para manejar métricas y resultados
// ---------------------------

// Metric representa una métrica con un autor, score y evidencias de violaciones DIP
type Metric struct {
	MetricID  string     `json:"metric_id"`
	GitAuthor string     `json:"git_author"`
	Score     int        `json:"score"`
	Evidence  []Evidence `json:"evidence"`
}

// Evidence representa evidencia de una violación de DIP, incluyendo el commit, archivo, línea y entidad afectada
type Evidence struct {
	CommitID   string `json:"commit_id"`
	File       string `json:"file"`
	Line       int    `json:"line"`
	EntityName string `json:"entity_name"`
}

// LayerConfig representa la configuración de capas desde el archivo YAML
type LayerConfig struct {
	Layers map[string][]string `yaml:"layers"`
}

// SkillData representa los datos de una habilidad, incluyendo el score y la evidencia encontrada
type SkillData struct {
	Score    int
	Evidence []Evidence
}

// Skill representa una habilidad a evaluar, como las violaciones DIP
type Skill struct {
	ID   string
	Name string
}
