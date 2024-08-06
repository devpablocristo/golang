package gmwsetup

import (
	"github.com/spf13/viper"
	"go-micro.dev/v4/registry"

	gmw "github.com/devpablocristo/golang/sdk/pkg/go-micro-web/v4"
)

func NewGoMicroInstance() (gmw.GoMicroClientPort, error) {
	config := gmw.GoMicroConfig{
		Name:     viper.GetString("MICRO_SERVICE_NAME"),
		Version:  viper.GetString("MICRO_SERVICE_VERSION"),
		Address:  viper.GetString("MICRO_SERVICE_ADDRESS"),
		Registry: registry.DefaultRegistry,
	}

	if err := gmw.InitializeGoMicroClient(config); err != nil {
		return nil, err
	}

	return gmw.GetGoMicroInstance()
}
