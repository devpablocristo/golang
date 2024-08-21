package hystrixconfig

import (
	"github.com/spf13/viper"

	htx "github.com/devpablocristo/golang/sdk/pkg/hystrix/hystrix-go"
)

func NewHystrixInstance() (htx.HystrixClientPort, error) {
	config := htx.HystrixConfig{
		Name:                   viper.GetString("HYSTRIX_NAME"),
		Timeout:                viper.GetInt("HYSTRIX_TIMEOUT"),
		MaxConcurrentRequests:  viper.GetInt("HYSTRIX_MAX_CONCURRENT_REQUESTS"),
		ErrorPercentThreshold:  viper.GetInt("HYSTRIX_ERROR_PERCENT_THRESHOLD"),
		RequestVolumeThreshold: viper.GetInt("HYSTRIX_REQUEST_VOLUME_THRESHOLD"),
		SleepWindow:            viper.GetInt("HYSTRIX_SLEEP_WINDOW"),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := htx.InitializeHystrixClient(config); err != nil {
		return nil, err
	}

	return htx.GetHystrixInstance()
}
