package cslhash

import (
	"fmt"
	"sync"

	"github.com/hashicorp/consul/api"
)

var (
	instance ConsulClientPort
	once     sync.Once
	errInit  error
)

type ConsulClient struct {
	client *api.Client
}

func InitializeConsulClient(config ConsulConfig) error {
	if err := config.Validate(); err != nil {
		return fmt.Errorf("invalid Consul configuration: %w", err)
	}

	once.Do(func() {
		client := &ConsulClient{}
		errInit = client.connect(config)
		if errInit != nil {
			instance = nil
		} else {
			instance = client
		}
	})
	return errInit
}

func GetConsulInstance() (ConsulClientPort, error) {
	if instance == nil {
		return nil, fmt.Errorf("consul client is not initialized")
	}
	return instance, nil
}

func (client *ConsulClient) connect(config ConsulConfig) error {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = config.Address

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to Consul: %w", err)
	}

	registration := &api.AgentServiceRegistration{
		ID:      config.ID,
		Name:    config.Name,
		Port:    config.Port,
		Address: config.Service,
		Check: &api.AgentServiceCheck{
			HTTP:     config.HealthCheck,
			Interval: config.CheckInterval,
			Timeout:  config.CheckTimeout,
		},
	}

	if err := consulClient.Agent().ServiceRegister(registration); err != nil {
		return err
	}

	client.client = consulClient
	return nil
}

func (client *ConsulClient) Client() *api.Client {
	return client.client
}
