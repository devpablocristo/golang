package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/teamcubation/monitors/commons-go/pkg/common"
	"github.com/teamcubation/monitors/commons-go/pkg/gitutils"
)

var skills = []common.Skill{
	{
		ID:    "go_unit_tests",
		Name:  "Writing unit tests using Go's testing package",
		Regex: regexp.MustCompile(`^func Test\w+\(t \*testing\.T\)`),
	},
	{
		ID:    "testify_assertions",
		Name:  "Using Testify for enhanced assertions",
		Regex: regexp.MustCompile(`github.com/stretchr/testify`),
	},
	{
		ID:    "table_driven_tests",
		Name:  "Implementing table-driven tests",
		Regex: regexp.MustCompile(`^\s*test\s*:=\s*\[\s*struct\s*{\s*name\s+string`),
		Checks: []func(string) bool{
			func(line string) bool {
				return regexp.MustCompile(`^\s*for\s+_,\s+tt\s+:=\s+range\s+test\s*{`).MatchString(line)
			},
		},
	},
	{
		ID:   "mock_objects",
		Name: "Using mock objects for testing",
		Checks: []func(string) bool{
			func(line string) bool { return strings.Contains(line, "github.com/golang/mock/gomock") },
			func(line string) bool { return strings.Contains(line, "go.uber.org/mock/gomock") },
			func(line string) bool { return strings.Contains(line, "github.com/stretchr/testify/mock") },
		},
	},
}

func analyzeRepo(repoPath string, files []string) ([]common.Metric, error) {
	repo, err := gitutils.GetRepo(repoPath)
	if err != nil {
		fmt.Printf("Error opening repository: %v\n", err)
		return nil, err
	}

	filesToAnalyze, err := gitutils.GetFilesToAnalyze(repo, files, func(path string) bool {
		return filepath.Ext(path) == ".go" && strings.HasSuffix(path, "_test.go")
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
				common.InitializeResultsForAuthor(skills, results, author)
			}

			common.UpdateResults(results[author], fileResults, filePath, commitID)
		}
	}

	return common.GenerateMetrics(skills, results), nil
}

func analyzeFile(filePath string) (map[string]common.SkillData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	results := common.InitializeResults(skills)

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		for _, skill := range skills {
			results[skill.ID] = common.AnalyzeSkill(skill, line, lineNumber, results[skill.ID])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
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
