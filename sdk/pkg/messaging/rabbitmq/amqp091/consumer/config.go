package consumer

import (
	"fmt"

	ports "github.com/devpablocristo/golang/sdk/pkg/messaging/rabbitmq/amqp091/consumer/ports"
)

type config struct {
	host     string
	port     int
	user     string
	password string
	vhost    string
}

func newConfig(host string, port int, user, password, vhost string) ports.Config {
	return &config{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		vhost:    vhost,
	}
}

func (c *config) GetHost() string     { return c.host }
func (c *config) SetHost(host string) { c.host = host }

func (c *config) GetPort() int     { return c.port }
func (c *config) SetPort(port int) { c.port = port }

func (c *config) GetUser() string     { return c.user }
func (c *config) SetUser(user string) { c.user = user }

func (c *config) GetPassword() string         { return c.password }
func (c *config) SetPassword(password string) { c.password = password }

func (c *config) GetVHost() string      { return c.vhost }
func (c *config) SetVHost(vhost string) { c.vhost = vhost }

func (c *config) Validate() error {
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
