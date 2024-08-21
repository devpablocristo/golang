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

// ConsulClientPort define la interfaz para interactuar con el cliente de Consul
type ConsulClientPort interface {
	Client() *api.Client
	Address() string // Añadir el método Address a la interfaz
}

// ConsulClient es la implementación del cliente de Consul
type ConsulClient struct {
	client  *api.Client
	address string // Almacenar la dirección aquí
}

// InitializeConsulClient inicializa el cliente de Consul
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

// GetConsulInstance devuelve la instancia de Consul
func GetConsulInstance() (ConsulClientPort, error) {
	if instance == nil {
		return nil, fmt.Errorf("consul client is not initialized")
	}
	return instance, nil
}

// connect conecta el cliente de Consul
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
		Tags:    config.Tags,
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
	client.address = consulConfig.Address // Almacenar la dirección

	return nil
}

// Client devuelve el cliente de Consul
func (client *ConsulClient) Client() *api.Client {
	return client.client
}

// Address devuelve la dirección del cliente de Consul
func (client *ConsulClient) Address() string {
	return client.address
}
