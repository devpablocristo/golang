package sdkgogit

import (
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

// Puedes añadir más métodos según tus necesidades
