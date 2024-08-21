package csdgocsl

import (
	"fmt"

	"github.com/devpablocristo/golang/sdk/pkg/cassandra/gocql/portspkg"
)

type cassandraConfig struct {
	hosts    []string
	keyspace string
	username string
	password string
}

func NewCassandraConfig(hosts []string, keyspace string, username string, password string) portspkg.CassandraConfig {
	h := make([]string, len(hosts))
	copy(h, hosts)
	return &cassandraConfig{
		hosts:    h,
		keyspace: keyspace,
		username: username,
		password: password,
	}
}

func (c *cassandraConfig) GetHosts() []string {
	return c.hosts
}

func (c *cassandraConfig) SetHosts(hosts []string) {
	h := make([]string, len(hosts))
	copy(h, hosts)
	c.hosts = h
}

func (c *cassandraConfig) GetKeyspace() string {
	return c.keyspace
}

func (c *cassandraConfig) SetKeyspace(keyspace string) {
	c.keyspace = keyspace
}

func (c *cassandraConfig) GetUsername() string {
	return c.username
}

func (c *cassandraConfig) SetUsername(username string) {
	c.username = username
}

func (c *cassandraConfig) GetPassword() string {
	return c.password
}

func (c *cassandraConfig) SetPassword(password string) {
	c.password = password
}

func (c *cassandraConfig) Validate() error {
	if len(c.hosts) == 0 {
		return fmt.Errorf("cassandra hosts are not configured")
	}
	if c.keyspace == "" {
		return fmt.Errorf("cassandra keyspace is not configured")
	}
	if c.username == "" {
		return fmt.Errorf("cassandra username is not configured")
	}
	if c.password == "" {
		return fmt.Errorf("cassandra password is not configured")
	}
	return nil
}
