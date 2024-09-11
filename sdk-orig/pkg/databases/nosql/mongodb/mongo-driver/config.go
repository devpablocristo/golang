package mongodbdriver

import (
	"fmt"
)

type config struct {
	User         string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

func newConfig(user, password, host, port, database string) *config {
	return &config{
		User:         user,
		Password:     password,
		Host:         host,
		Port:         port,
		DatabaseName: database, // Renombrado
	}
}

func (c *config) DSN() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		c.User, c.Password, c.Host, c.Port, c.DatabaseName)
}

func (c *config) Database() string {
	return c.DatabaseName
}

func (c *config) Validate() error {
	if c.User == "" || c.Password == "" || c.Host == "" || c.Port == "" || c.DatabaseName == "" {
		return fmt.Errorf("incomplete MongoDB configuration")
	}
	return nil
}
