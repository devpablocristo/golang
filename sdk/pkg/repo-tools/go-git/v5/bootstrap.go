package sdkgogit

import (
	ports "github.com/devpablocristo/golang/sdk/pkg/repo-tools/go-git/v5/ports"
	"github.com/spf13/viper"
)

func Bootstrap(repoRemoteUrlKey, repoLocalPathKey, repoBranchKey string) (ports.Client, error) {
	config := newConfig(
		viper.GetString(repoRemoteUrlKey),
		viper.GetString(repoLocalPathKey),
		viper.GetString(repoBranchKey),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newClient(config)
}
