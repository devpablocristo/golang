package cslhash

import (
	"fmt"
)

type ConsulConfig struct {
	ID            string
	Name          string
	Port          int
	Address       string
	Service       string
	HealthCheck   string
	CheckInterval string
	CheckTimeout  string
	Tags          []string
}

func (c ConsulConfig) Validate() error {
	if c.ID == "" || c.Name == "" || c.Port == 0 || c.Address == "" || c.HealthCheck == "" || c.Service == "" {
		return fmt.Errorf("incomplete Consul configuration")
	}
	return nil
}
