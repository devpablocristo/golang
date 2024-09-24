package ports

type Server interface {
	Run() error
	SetWebRouter(router interface{}) error
}

type Config interface {
	GetServerName() string
	GetServerHost() string
	GetServerPort() int
	GetServerAddress() string
	GetConsulAddress() string
	GetRouter() interface{}
	Validate() error
}
