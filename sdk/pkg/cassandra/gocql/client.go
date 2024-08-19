package csdgocsl

import (
	"fmt"
	"sync"

	"github.com/gocql/gocql"

	"github.com/devpablocristo/golang/sdk/pkg/cassandra/gocql/portspkg"
)

var (
	instance portspkg.CassandraClient
	once     sync.Once
	errInit  error
)

type cassandraClient struct {
	session *gocql.Session
}

func InitializeCassandraClient(config portspkg.CassandraConfig) error {
	once.Do(func() {
		client := &cassandraClient{}
		errInit = client.Connect(config)
		if errInit == nil {
			instance = client
		}
	})
	return errInit
}

func GetCassandraInstance() (portspkg.CassandraClient, error) {
	if instance == nil {
		return nil, fmt.Errorf("cassandra client is not initialized")
	}
	return instance, nil
}

func (c *cassandraClient) Connect(config portspkg.CassandraConfig) error {
	cluster := gocql.NewCluster(config.GetHosts()...)
	cluster.Keyspace = config.GetKeyspace()
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: config.GetUsername(),
		Password: config.GetPassword(),
	}
	session, err := cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("failed to connect to Cassandra: %w", err)
	}
	c.session = session
	return nil
}

func (c *cassandraClient) Close() {
	if c.session != nil {
		c.session.Close()
	}
}

func (c *cassandraClient) GetSession() *gocql.Session {
	return c.session
}
