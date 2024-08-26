package producer

import (
	"fmt"

	ports "github.com/devpablocristo/golang/sdk/pkg/messaging/rabbitmq/amqp091/producer/ports"
)

// producerConfig estructura que implementa la interfaz ports.Config para el productor
type producerConfig struct {
	host     string
	port     int
	user     string
	password string
	vhost    string
}

// newProducerConfig crea una nueva configuración para el productor de RabbitMQ
func newConfig(host string, port int, user, password, vhost string) ports.Config {
	return &producerConfig{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		vhost:    vhost,
	}
}

func (c *producerConfig) GetHost() string     { return c.host }
func (c *producerConfig) SetHost(host string) { c.host = host }

func (c *producerConfig) GetPort() int     { return c.port }
func (c *producerConfig) SetPort(port int) { c.port = port }

func (c *producerConfig) GetUser() string     { return c.user }
func (c *producerConfig) SetUser(user string) { c.user = user }

func (c *producerConfig) GetPassword() string         { return c.password }
func (c *producerConfig) SetPassword(password string) { c.password = password }

func (c *producerConfig) GetVHost() string      { return c.vhost }
func (c *producerConfig) SetVHost(vhost string) { c.vhost = vhost }

func (c *producerConfig) Validate() error {
	if c.host == "" {
		return fmt.Errorf("rabbitmq host is not configured")
	}
	if c.port == 0 {
		return fmt.Errorf("rabbitmq port is not configured")
	}
	if c.user == "" {
		return fmt.Errorf("rabbitmq user is not configured")
	}
	if c.password == "" {
		return fmt.Errorf("rabbitmq password is not configured")
	}
	if c.vhost == "" {
		return fmt.Errorf("rabbitmq vhost is not configured")
	}
	return nil
}
