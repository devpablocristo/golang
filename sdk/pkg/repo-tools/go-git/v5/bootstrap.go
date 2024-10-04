package sdkgogit

import (
	ports "github.com/devpablocristo/golang/sdk/pkg/repo-tools/go-git/v5/ports"
)

func Bootstrap() (ports.GitClient, error) {
	// config := newConfig(
	// 	viper.GetString("GIT_REPO_URL"),
	// 	viper.GetString("GIT_REPO_PATH"),
	// 	viper.GetString("GIT_REPO_BRANCH"),
	// )

	config := newConfig(
		"https://github.com/devpablocristo/meli",
		"/home/pablo/tests",
		"main",
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newClient(config)
}
