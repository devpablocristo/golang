package common

import "regexp"

type Metric struct {
	MetricID  string     `json:"metric_id"`
	GitAuthor string     `json:"git_author"`
	Score     string     `json:"score"`
	Evidence  []Evidence `json:"evidence"`
}

type Evidence struct {
	CommitID string `json:"commit_id"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

type SkillData struct {
	Score    int
	Evidence []Evidence
}

type Skill struct {
	ID     string
	Name   string
	Regex  *regexp.Regexp
	Checks []func(string) bool
}

type Results map[string]map[string]SkillData
