package ctypes

type ConfigMongoPort interface {
	GetMongoURL() string
	GetMongoDBName() string
	GetMongoCollectionName() string
}

type ConfigGinPort interface {
	GetHandlerPort() string
}
