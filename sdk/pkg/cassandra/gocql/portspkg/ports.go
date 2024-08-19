package portspkg

import (
	"github.com/gocql/gocql"
)

type CassandraConfig interface {
	GetHosts() []string
	SetHosts(hosts []string)
	GetKeyspace() string
	SetKeyspace(keyspace string)
	GetUsername() string
	SetUsername(username string)
	GetPassword() string
	SetPassword(password string)
	Validate() error
}

type CassandraClient interface {
	Connect(config CassandraConfig) error
	Close()
	GetSession() *gocql.Session
}
