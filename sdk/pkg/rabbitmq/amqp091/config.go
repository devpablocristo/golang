package amsgqp

import "fmt"

// RabbitMQConfig representa la configuración necesaria para conectarse a RabbitMQ.
type RabbitMQConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	VHost    string
}

// Validate valida la configuración de RabbitMQ.
func (config *RabbitMQConfig) Validate() error {
	if config.Host == "" {
		return fmt.Errorf("rabbitmq host is not configured")
	}
	if config.Port == 0 {
		return fmt.Errorf("rabbitmq port is not configured")
	}
	if config.User == "" {
		return fmt.Errorf("rabbitmq user is not configured")
	}
	if config.Password == "" {
		return fmt.Errorf("rabbitmq password is not configured")
	}
	if config.VHost == "" {
		return fmt.Errorf("rabbitmq vhost is not configured")
	}
	return nil
}
