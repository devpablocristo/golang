package ports

import (
	"github.com/go-git/go-git/v5"
)

type GitClient interface {
	GetRepository() *git.Repository
	PullLatest() error
	GetFilesToAnalyze([]string, func(string) bool) ([]string, error)
	GetFileAuthor(string) (string, error)
	GetCommitID(string) (string, error)
	GetRepo(string) (*git.Repository, error)
}

type Config interface {
	GetRepoURL() string
	SetRepoURL(string)
	GetRepoPath() string
	SetRepoPath(string)
	GetRepoBranch() string
	SetRepoBranch(string)
	Validate() error
}
