package pyrosgo

import (
	"fmt"
	"log"
	"sync"

	"github.com/grafana/pyroscope-go"
)

var (
	instance PyroscopeClientPort
	once     sync.Once
	errInit  error
)

type PyroscopeClient struct{}

func InitializePyroscopeClient(config PyroscopeClientConfig) error {
	once.Do(func() {
		client := &PyroscopeClient{}
		errInit = client.connect(config)
		if errInit != nil {
			instance = nil
		} else {
			instance = client
		}
	})
	return errInit
}

func GetPyroscopeInstance() (PyroscopeClientPort, error) {
	if instance == nil {
		return nil, fmt.Errorf("client Pyroscope: client is not initialized")
	}
	return instance, nil
}

func (client *PyroscopeClient) connect(config PyroscopeClientConfig) error {
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: config.ApplicationName,
		ServerAddress:   config.ServerAddress,
		AuthToken:       config.AuthToken,
		ProfileTypes:    []pyroscope.ProfileType{pyroscope.ProfileCPU, pyroscope.ProfileAllocObjects, pyroscope.ProfileAllocSpace, pyroscope.ProfileInuseObjects, pyroscope.ProfileInuseSpace},
	})
	if err != nil {
		return fmt.Errorf("failed to start Pyroscope: %w", err)
	}
	log.Println("Pyroscope started successfully")
	return nil
}

func (client *PyroscopeClient) Close() {
	// Pyroscope no tiene un método Close explícito, pero puedes limpiar recursos aquí si es necesario
}
