package csdgocsl

import (
	"fmt"
	"sync"

	"github.com/gocql/gocql"
)

var (
	instance CassandraClientPort
	once     sync.Once
	errInit  error
)

type CassandraClient struct {
	session *gocql.Session
}

func InitializeCassandraClient(config CassandraConfig) error {
	once.Do(func() {
		instance = &CassandraClient{}
		errInit = instance.Connect(config)
		if errInit != nil {
			instance = nil
		}
	})
	return errInit
}

func GetCassandraInstance() (CassandraClientPort, error) {
	if instance == nil {
		return nil, fmt.Errorf("cassandra client is not initialized")
	}
	return instance, nil
}

func (c *CassandraClient) Connect(config CassandraConfig) error {
	cluster := gocql.NewCluster(config.Hosts...)
	cluster.Keyspace = config.Keyspace
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: config.Username,
		Password: config.Password,
	}
	session, err := cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("failed to connect to Cassandra: %w", err)
	}
	c.session = session
	return nil
}

func (c *CassandraClient) Close() {
	if c.session != nil {
		c.session.Close()
	}
}

func (c *CassandraClient) GetSession() *gocql.Session {
	return c.session
}
