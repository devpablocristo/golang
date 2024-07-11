package csdgocsl

import (
	"github.com/gocql/gocql"
)

type CassandraClientPort interface {
	Connect(config CassandraConfig) error
	Close()
	GetSession() *gocql.Session
}
