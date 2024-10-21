package common

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/teamcubation/monitors/commons-go/pkg/gitutils"
)

func InitializeResults(skills []Skill) map[string]SkillData {
	results := make(map[string]SkillData)
	for _, skill := range skills {
		results[skill.ID] = SkillData{Score: 0, Evidence: []Evidence{}}
	}
	return results
}

func InitializeResultsForAuthor(skills []Skill, results map[string]map[string]SkillData, author string) {
	results[author] = make(map[string]SkillData)
	for _, skill := range skills {
		results[author][skill.ID] = SkillData{Score: 0, Evidence: []Evidence{}}
	}
}

func UpdateResults(authorResults map[string]SkillData, fileResults map[string]SkillData, filePath, commitID string) {
	for skillID, data := range fileResults {
		skillData := authorResults[skillID]
		if data.Score > skillData.Score {
			skillData.Score = data.Score
			skillData.Evidence = []Evidence{}
		}
		if data.Score == skillData.Score {
			addEvidenceToSkillData(&skillData, data.Evidence, filePath, commitID)
		}
		authorResults[skillID] = skillData
	}
}

func addEvidenceToSkillData(skillData *SkillData, evidence []Evidence, filePath, commitID string) {
	for _, line := range evidence {
		skillData.Evidence = append(skillData.Evidence, Evidence{
			CommitID: commitID,
			File:     filePath,
			Line:     line.Line,
		})
	}
}

func GenerateMetrics(skills []Skill, results map[string]map[string]SkillData) []Metric {
	var output []Metric
	for author, scores := range results {
		for _, skill := range skills {
			skillID := skill.ID
			output = append(output, Metric{
				MetricID:  skillID,
				GitAuthor: author,
				Score:     fmt.Sprintf("%d", scores[skillID].Score),
				Evidence:  scores[skillID].Evidence,
			})
		}
	}
	return output
}

func AnalyzeSkill(skill Skill, line string, lineNumber int, skillData SkillData) SkillData {
	if skill.Regex != nil && skill.Regex.MatchString(line) {
		skillData.Score = 1
		skillData.Evidence = append(skillData.Evidence, Evidence{Line: lineNumber})
	}

	for _, check := range skill.Checks {
		if check(line) {
			skillData.Score = 1
			skillData.Evidence = append(skillData.Evidence, Evidence{Line: lineNumber})
		}
	}
	return skillData
}

func ShouldAnalyzeFile(fullPath, filePath string) bool {
	fileInfo, err := os.Stat(fullPath)
	return err == nil && !fileInfo.IsDir() && filepath.Ext(filePath) == ".go"
}

func GetAuthorAndCommit(repo *git.Repository, filePath string) (string, string, error) {
	author, err := gitutils.GetFileAuthor(repo, filePath)
	if err != nil {
		return "", "", err
	}
	commitID, err := gitutils.GetCommitID(repo, filePath)
	if err != nil {
		return "", "", err
	}
	return author, commitID, nil
}
