package pkgcassandra

import (
	"fmt"
	"sync"

	"github.com/gocql/gocql"

	ports "github.com/devpablocristo/golang/sdk/pkg/databases/nosql/cassandra/gocql/ports"
)

var (
	instance  ports.Service
	once      sync.Once
	initError error
)

type service struct {
	session *gocql.Session
}

func NewService(config ports.Config) (ports.Service, error) {
	once.Do(func() {
		client := &service{}
		initError = client.Connect(config)
		if initError == nil {
			instance = client
		}
	})
	return instance, initError
}

func GetInstance() (ports.Service, error) {
	if instance == nil {
		return nil, fmt.Errorf("cassandra client is not initialized")
	}
	return instance, nil
}

func (c *service) Connect(config ports.Config) error {
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

func (c *service) Close() {
	if c.session != nil {
		c.session.Close()
	}
}

func (c *service) GetSession() *gocql.Session {
	return c.session
}
