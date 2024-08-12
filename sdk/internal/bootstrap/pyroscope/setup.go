package pyroscopesetup

import (
	"github.com/spf13/viper"

	pyrosgo "github.com/devpablocristo/golang/sdk/pkg/pyroscope/pyroscope-go"
)

func NewPyroscopeInstance() (pyrosgo.PyroscopeClientPort, error) {
	config := pyrosgo.PyroscopeClientConfig{
		ApplicationName: viper.GetString("PYROSCOPE_APPLICATION_NAME"),
		ServerAddress:   viper.GetString("PYROSCOPE_SERVER_ADDRESS"),
		AuthToken:       viper.GetString("PYROSCOPE_AUTH_TOKEN"),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := pyrosgo.InitializePyroscopeClient(config); err != nil {
		return nil, err
	}

	return pyrosgo.GetPyroscopeInstance()
}
