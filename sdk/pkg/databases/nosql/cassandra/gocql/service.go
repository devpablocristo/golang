package sdkcassandra

import (
	"fmt"
	"sync"

	"github.com/gocql/gocql"

	ports "github.com/devpablocristo/golang/sdk/pkg/databases/nosql/cassandra/gocql/ports"
)

var (
	instance  ports.Repository
	once      sync.Once
	initError error
)

type Repository struct {
	session *gocql.Session
}

func NewRepository(config ports.Config) (ports.Repository, error) {
	once.Do(func() {
		client := &Repository{}
		initError = client.Connect(config)
		if initError == nil {
			instance = client
		}
	})
	return instance, initError
}

func GetInstance() (ports.Repository, error) {
	if instance == nil {
		return nil, fmt.Errorf("cassandra client is not initialized")
	}
	return instance, nil
}

func (c *Repository) Connect(config ports.Config) error {
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

func (c *Repository) Close() {
	if c.session != nil {
		c.session.Close()
	}
}

func (c *Repository) GetSession() *gocql.Session {
	return c.session
}
