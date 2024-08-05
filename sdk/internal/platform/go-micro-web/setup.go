package gmwsetup

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/spf13/viper"
	"go-micro.dev/v4/registry"

	cslhash "github.com/devpablocristo/golang/sdk/pkg/consul/hashicorp"
	gmw "github.com/devpablocristo/golang/sdk/pkg/go-micro-web"
)

func NewGoMicroInstance(consulInst cslhash.ConsulClientPort) (gmw.GoMicroClientPort, error) {
	consulAdreess := viper.GetString("CONSUL_ADDRESS")
	if consulInst != nil {
		consulAdreess = consulInst.Address()
	}
	consulRegistry := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{consulAdreess}
	})

	config := gmw.GoMicroConfig{
		Name:    viper.GetString("MICRO_SERVICE_NAME"),
		Version: viper.GetString("MICRO_SERVICE_VERSION"),
		Address: viper.GetString("MICRO_SERVICE_ADDRESS"),
		//Registry: registry.DefaultRegistry,
		Registry: consulRegistry,
	}

	if err := gmw.InitializeGoMicroClient(config); err != nil {
		return nil, err
	}

	return gmw.GetGoMicroInstance()
}
