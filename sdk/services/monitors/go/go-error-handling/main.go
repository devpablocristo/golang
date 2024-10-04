package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/teamcubation/monitors/commons-go/pkg/common"
	"github.com/teamcubation/monitors/commons-go/pkg/gitutils"
)

var errorHandlingSkills = []common.Skill{
	{
		ID:    "error_checking",
		Name:  "Verifying return values of functions that return error",
		Regex: regexp.MustCompile(`if\s+err\s*!=\s*nil`),
	},
	{
		ID:   "panic_usage",
		Name: "Minimal and correct usage of panic/recover",
		Regex: regexp.MustCompile(strings.Join([]string{
			`panic\(`,
			`log\.Fatal\(`,
			`log\.Fatalf\(`,
			`os\.Exit\(`,
		}, "|")),
	},
	{
		ID:    "error_context",
		Name:  "Providing additional context in error messages (using fmt.Errorf or pkg/errors)",
		Regex: regexp.MustCompile(`fmt\.Errorf\(|errors\.Wrap\(`),
	},
	{ID: "ignored_errors", Name: "Minimizing ignored errors"},
	{ID: "defer_usage", Name: "Using defer to manage resource cleanup in case of errors"},
	{ID: "log_error_count", Name: "Monitoring and minimizing frequent error logs during runtime"},
}

func analyzeRepo(repoPath string, files []string) ([]common.Metric, error) {
	repo, err := gitutils.GetRepo(repoPath)
	if err != nil {
		fmt.Printf("Error opening repository: %v\n", err)
		return nil, err
	}

	filesToAnalyze, err := gitutils.GetFilesToAnalyze(repo, files, func(path string) bool {
		return filepath.Ext(path) == ".go"
	})
	if err != nil {
		return nil, err
	}

	results := make(map[string]map[string]common.SkillData)

	for _, filePath := range filesToAnalyze {
		fullPath := filepath.Join(repoPath, filePath)
		if common.ShouldAnalyzeFile(fullPath, filePath) {
			author, commitID, err := common.GetAuthorAndCommit(repo, filePath)
			if err != nil {
				return nil, err
			}

			fileResults, err := analyzeFile(fullPath)
			if err != nil {
				return nil, err
			}

			if _, exists := results[author]; !exists {
				common.InitializeResultsForAuthor(errorHandlingSkills, results, author)
			}

			common.UpdateResults(results[author], fileResults, filePath, commitID)
		}
	}

	return common.GenerateMetrics(errorHandlingSkills, results), nil
}

func analyzeFile(filePath string) (map[string]common.SkillData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	results := common.InitializeResults(errorHandlingSkills)

	checkIgnoredErrors(filePath, results)

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	errorCheckRegex := regexp.MustCompile(`if\s+err\s*!=\s*nil`)
	errorContextRegex := regexp.MustCompile(`fmt\.Errorf\(|errors\.Wrap\(`)
	// ignoredErrorRegex := regexp.MustCompile(`^[^if]*err\s*=\s*[^\n]+$`)

	patterns := []string{
		`panic\(`,
		`log\.Fatal\(`,
		`log\.Fatalf\(`,
		`os\.Exit\(`,
	}

	panicUsageRegex := regexp.MustCompile(strings.Join(patterns, "|"))

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		for _, skill := range errorHandlingSkills {
			skillData := results[skill.ID]
			switch skill.ID {
			case "error_checking":
				if errorCheckRegex.MatchString(line) {
					skillData.Score = 1
					skillData.Evidence = append(skillData.Evidence, common.Evidence{Line: lineNumber})
				}
			case "panic_usage":
				if strings.Contains(filePath, "cmd") {
					continue
				}
				if panicUsageRegex.MatchString(line) {
					skillData.Score = 0
					skillData.Evidence = append(skillData.Evidence, common.Evidence{Line: lineNumber})
				}
			case "error_context":
				if errorContextRegex.MatchString(line) {
					skillData.Score = 1
					skillData.Evidence = append(skillData.Evidence, common.Evidence{Line: lineNumber})
				}
			case "ignored_errors":
				continue
			}
			results[skill.ID] = skillData
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func inspectAST(f *ast.File, fset *token.FileSet, results map[string]common.SkillData) {
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.AssignStmt:
			for i, lhs := range x.Lhs {
				if ident, ok := lhs.(*ast.Ident); ok && ident.Name == "_" {
					if i >= len(x.Rhs) {
						skillData := results["ignored_errors"]
						skillData.Score = 0
						skillData.Evidence = append(skillData.Evidence, common.Evidence{Line: fset.Position(ident.Pos()).Line})
						results["ignored_errors"] = skillData
					}
				}
			}
		}
		return true
	})
}

func checkIgnoredErrors(filePath string, results map[string]common.SkillData) {
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	inspectAST(file, fset, results)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <repo_path> [file1] [file2] ...")
		return
	}

	repoPath := os.Args[1]
	filesToAnalyze := os.Args[2:]

	metrics, err := analyzeRepo(repoPath, filesToAnalyze)
	if err != nil {
		log.Fatal(err)
		return
	}

	jsonOutput, err := json.MarshalIndent(metrics, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonOutput))
}
