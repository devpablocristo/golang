package sdkgogit

import (
	ports "github.com/devpablocristo/golang/sdk/pkg/repo-tools/go-git/v5/ports"
	"github.com/spf13/viper"
)

func Bootstrap(repoRemoteUrlEnvName, repoLocalPathEnvName, repoBranchEnvName string) (ports.Client, error) {
	config := newConfig(
		viper.GetString(repoRemoteUrlEnvName),
		viper.GetString(repoLocalPathEnvName),
		viper.GetString(repoBranchEnvName),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newClient(config)
}
