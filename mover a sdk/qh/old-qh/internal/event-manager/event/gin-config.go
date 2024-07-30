package event

import (
	"os"
)

type ConfigGin struct {
	Port string
}

func NewGinConfig() ConfigGinPort {
	return &ConfigGin{
		Port: os.Getenv("HANDLER_PORT"),
	}
}

func (c *ConfigGin) GetHandlerPort() string {
	return c.Port
}
