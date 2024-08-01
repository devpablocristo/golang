package gmwsetup

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/spf13/viper"
	"go-micro.dev/v4/registry"

	gmw "github.com/devpablocristo/qh/events/pkg/go-micro-web"
)

func NewGoMicroInstance() (gmw.GoMicroClientPort, error) {

	ref := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{viper.GetString("CONSUL_ADDRESS")}
	})

	config := gmw.GoMicroConfig{
		Name:    viper.GetString("MICRO_SERVICE_NAME"),
		Version: viper.GetString("MICRO_SERVICE_VERSION"),
		Address: viper.GetString("MICRO_SERVICE_ADDRESS"),
		//Registry: registry.DefaultRegistry,
		Registry: ref,
	}

	if err := gmw.InitializeGoMicroClient(config); err != nil {
		return nil, err
	}

	return gmw.GetGoMicroInstance()
}
