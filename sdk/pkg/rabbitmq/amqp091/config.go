package amsgqp

import (
	"fmt"

	"github.com/devpablocristo/golang/sdk/pkg/rabbitmq/amqp091/pkgports"
)

// rabbitMqConfig representa la configuración necesaria para conectarse a RabbitMQ.
type rabbitMqConfig struct {
	host     string
	port     int
	user     string
	password string
	vhost    string
}

// NewRabbitMqConfig crea una nueva configuración de RabbitMQ con los valores proporcionados.
func NewRabbitMqConfig(host string, port int, user, password, vhost string) pkgports.RabbitMqConfig {
	return &rabbitMqConfig{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		vhost:    vhost,
	}
}

// GetHost devuelve el host de la configuración.
func (config *rabbitMqConfig) GetHost() string {
	return config.host
}

// SetHost establece el host de la configuración.
func (config *rabbitMqConfig) SetHost(host string) {
	config.host = host
}

// GetPort devuelve el puerto de la configuración.
func (config *rabbitMqConfig) GetPort() int {
	return config.port
}

// SetPort establece el puerto de la configuración.
func (config *rabbitMqConfig) SetPort(port int) {
	config.port = port
}

// GetUser devuelve el usuario de la configuración.
func (config *rabbitMqConfig) GetUser() string {
	return config.user
}

// SetUser establece el usuario de la configuración.
func (config *rabbitMqConfig) SetUser(user string) {
	config.user = user
}

// GetPassword devuelve la contraseña de la configuración.
func (config *rabbitMqConfig) GetPassword() string {
	return config.password
}

// SetPassword establece la contraseña de la configuración.
func (config *rabbitMqConfig) SetPassword(password string) {
	config.password = password
}

// GetVHost devuelve el Virtual Host de la configuración.
func (config *rabbitMqConfig) GetVHost() string {
	return config.vhost
}

// SetVHost establece el Virtual Host de la configuración.
func (config *rabbitMqConfig) SetVHost(vhost string) {
	config.vhost = vhost
}

// Validate valida la configuración de RabbitMQ.
func (config *rabbitMqConfig) Validate() error {
	if config.host == "" {
		return fmt.Errorf("rabbitmq host is not configured")
	}
	if config.port == 0 {
		return fmt.Errorf("rabbitmq port is not configured")
	}
	if config.user == "" {
		return fmt.Errorf("rabbitmq user is not configured")
	}
	if config.password == "" {
		return fmt.Errorf("rabbitmq password is not configured")
	}
	if config.vhost == "" {
		return fmt.Errorf("rabbitmq vhost is not configured")
	}
	return nil
}
