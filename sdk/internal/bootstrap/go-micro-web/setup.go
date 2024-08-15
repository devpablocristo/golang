package gmwsetup

import (
	"github.com/spf13/viper"
	"go-micro.dev/v4/registry"

	gmw "github.com/devpablocristo/golang/sdk/pkg/go-micro-web/v4"
	portspkg "github.com/devpablocristo/golang/sdk/pkg/go-micro-web/v4/portspkg"
)

func NewGoMicroInstance() (portspkg.GoMicroClient, error) {
	config := gmw.NewGoMicroConfig(
		viper.GetString("MICRO_SERVICE_NAME"),
		viper.GetString("MICRO_SERVICE_VERSION"),
		viper.GetString("MICRO_SERVICE_ADDRESS"),
		registry.DefaultRegistry,
	)

	if err := gmw.InitializeGoMicroClient(config); err != nil {
		return nil, err
	}

	return gmw.GetGoMicroClientInstance()
}
