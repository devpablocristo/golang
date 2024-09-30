package ports

import "go-micro.dev/v4/client"

type Client interface {
	GetClient() client.Client
}

type Config interface {
	Validate() error
	GetConsulAddress() string
	GetServerName() string
}
