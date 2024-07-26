package hystrixgo

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/afex/hystrix-go/hystrix"
)

var (
	instance HystrixClientPort
	once     sync.Once
	errInit  error
)

type HystrixClient struct {
	config *HystrixConfig
}

func InitializeHystrixClient(config HystrixConfig) error {
	if err := config.Validate(); err != nil {
		return fmt.Errorf("invalid Hystrix configuration: %w", err)
	}

	once.Do(func() {
		client := &HystrixClient{config: &config}
		errInit = client.connect(config)
		if errInit != nil {
			instance = nil
		} else {
			instance = client
		}
	})
	return errInit
}

func GetHystrixInstance() (HystrixClientPort, error) {
	if instance == nil {
		return nil, fmt.Errorf("Hystrix client is not initialized")
	}
	return instance, nil
}

func (client *HystrixClient) connect(config HystrixConfig) error {
	hystrix.ConfigureCommand(config.Name, hystrix.CommandConfig{
		Timeout:                config.Timeout,
		MaxConcurrentRequests:  config.MaxConcurrentRequests,
		ErrorPercentThreshold:  config.ErrorPercentThreshold,
		RequestVolumeThreshold: config.RequestVolumeThreshold,
		SleepWindow:            config.SleepWindow,
	})
	return nil
}

func (client *HystrixClient) Get(url string) (*http.Response, error) {
	var resp *http.Response
	err := hystrix.Do(client.config.Name, func() error {
		var err error
		resp, err = http.Get(url)
		if err != nil {
			return err
		}
		return nil
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get: %w", err)
	}
	return resp, nil
}
