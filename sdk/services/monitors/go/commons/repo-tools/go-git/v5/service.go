package sdkgogit

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	ports "github.com/devpablocristo/golang/sdk/pkg/repo-tools/go-git/v5/ports"
)

var (
	instance  ports.GitClient
	once      sync.Once
	initError error
)

type client struct {
	repo   *git.Repository
	config ports.Config
}

func newClient(config ports.Config) (ports.GitClient, error) {
	once.Do(func() {
		err := config.Validate()
		if err != nil {
			initError = err
			return
		}

		// Clonar o abrir el repositorio
		repo, err := git.PlainOpen(config.GetRepoPath())
		if err != nil {
			repo, err = git.PlainClone(config.GetRepoPath(), false, &git.CloneOptions{
				URL:           config.GetRepoURL(),
				ReferenceName: plumbing.NewBranchReferenceName(config.GetRepoBranch()),
				Progress:      nil,
			})
			if err != nil {
				initError = err
				return
			}
		}

		instance = &client{
			config: config,
			repo:   repo,
		}
	})
	return instance, initError
}

func (gc *client) GetRepository() *git.Repository {
	return gc.repo
}

func (gc *client) PullLatest() error {
	w, err := gc.repo.Worktree()
	if err != nil {
		return err
	}

	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}
	return nil
}

func (gc *client) GetFilesToAnalyze(files []string, filterFunc func(string) bool) ([]string, error) {
	var filesToAnalyze []string

	if len(files) == 0 {
		worktree, err := gc.repo.Worktree()
		if err != nil {
			return nil, err
		}

		err = filepath.Walk(worktree.Filesystem.Root(), func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && filterFunc(path) {
				relPath, err := filepath.Rel(worktree.Filesystem.Root(), path)
				if err != nil {
					return err
				}
				filesToAnalyze = append(filesToAnalyze, relPath)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		for _, file := range files {
			if filterFunc(file) {
				filesToAnalyze = append(filesToAnalyze, file)
			}
		}
	}

	return filesToAnalyze, nil
}

func (gc *client) GetFileAuthor(file string) (string, error) {
	commits, err := gc.repo.Log(&git.LogOptions{FileName: &file})
	if err != nil {
		return "", err
	}

	commit, err := commits.Next()
	if err != nil {
		return "", err
	}

	return commit.Author.Email, nil
}

// GetCommitID retorna el ID del último commit de un archivo
func (gc *client) GetCommitID(file string) (string, error) {
	commits, err := gc.repo.Log(&git.LogOptions{FileName: &file})
	if err != nil {
		return "", err
	}

	objectCommit, err := commits.Next()
	if err != nil {
		return "", err
	}

	return objectCommit.Hash.String(), nil
}

func (gc *client) GetRepo(repoPath string) (*git.Repository, error) {
	return git.PlainOpen(repoPath)
}
	