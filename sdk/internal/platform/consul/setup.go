package consulsetup

import (
	"github.com/spf13/viper"

	cslhash "github.com/devpablocristo/golang/sdk/pkg/consul/hashicorp"
)

func NewConsulInstance() (cslhash.ConsulClientPort, error) {
	config := cslhash.ConsulConfig{
		ID:            viper.GetString("CONSUL_ID"),
		Name:          viper.GetString("CONSUL_NAME"),
		Port:          viper.GetInt("CONSUL_PORT"),
		Address:       viper.GetString("CONSUL_ADDRESS"),
		Service:       viper.GetString("CONSUL_SERVICE_NAME"),
		HealthCheck:   viper.GetString("CONSUL_HEALTH_CHECK"),
		CheckInterval: viper.GetString("CONSUL_CHECK_INTERVAL"),
		CheckTimeout:  viper.GetString("CONSUL_CHECK_TIMEOUT"),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := cslhash.InitializeConsulClient(config); err != nil {
		return nil, err
	}

	return cslhash.GetConsulInstance()
}
