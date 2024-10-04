package ports

import (
	"github.com/go-git/go-git/v5"
)

type GitClient interface {
	GetRepository() *git.Repository
	PullLatest() error
	// Añade más métodos según tus necesidades
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
