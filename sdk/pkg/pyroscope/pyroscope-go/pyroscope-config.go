package pyrosgo

import "fmt"

type PyroscopeClientConfig struct {
	ApplicationName string
	ServerAddress   string
	AuthToken       string // Si tu servidor Pyroscope requiere autenticación
}

func (config PyroscopeClientConfig) Validate() error {
	if config.ApplicationName == "" || config.ServerAddress == "" {
		return fmt.Errorf("incomplete Pyroscope configuration")
	}
	return nil
}
