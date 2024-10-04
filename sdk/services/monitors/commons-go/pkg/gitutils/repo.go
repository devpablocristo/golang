package gitutils

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func GetFilesToAnalyze(repo *git.Repository, files []string, filterFunc func(string) bool) ([]string, error) {
	var filesToAnalyze []string

	if len(files) == 0 {
		worktree, err := repo.Worktree()
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

func GetFileAuthor(repo *git.Repository, file string) (string, error) {
	commits, err := repo.Log(&git.LogOptions{FileName: &file})
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
func GetCommitID(repo *git.Repository, file string) (string, error) {
	commits, err := repo.Log(&git.LogOptions{FileName: &file})
	if err != nil {
		return "", err
	}

	objectCommit, err := commits.Next()
	if err != nil {
		return "", err
	}

	return objectCommit.Hash.String(), nil
}

func GetRepo(repoPath string) (*git.Repository, error) {
	return git.PlainOpen(repoPath)
}
