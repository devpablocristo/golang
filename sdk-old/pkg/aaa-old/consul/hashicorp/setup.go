package consulsetup

import (
	"strings"

	"github.com/spf13/viper"

	sdkconsul "github.com/devpablocristo/golang/sdk/pkg/consul/hashicorp"
)

func NewConsulInstance() (sdkconsul.ConsulClientPort, error) {
	tagsString := viper.GetString("CONSUL_TAGS")
	tags := strings.Split(tagsString, ",") // Asume que los tags están separados por comas

	config := sdkconsul.ConsulConfig{
		ID:            viper.GetString("CONSUL_ID"),
		Name:          viper.GetString("CONSUL_NAME"),
		Port:          viper.GetInt("CONSUL_PORT"),
		Address:       viper.GetString("CONSUL_ADDRESS"),
		Service:       viper.GetString("CONSUL_SERVICE_NAME"),
		HealthCheck:   viper.GetString("CONSUL_HEALTH_CHECK"),
		CheckInterval: viper.GetString("CONSUL_CHECK_INTERVAL"),
		CheckTimeout:  viper.GetString("CONSUL_CHECK_TIMEOUT"),
		Tags:          tags,
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := sdkconsul.InitializeConsulClient(config); err != nil {
		return nil, err
	}

	return sdkconsul.GetConsulInstance()
}
