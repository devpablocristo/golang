package mongodbdriver

import "fmt"

// MongoDBClientConfig representa a configuração necessária para conectar ao MongoDB
type MongoDBClientConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

// dsn retorna a string de conexão para o MongoDB
func (config MongoDBClientConfig) dsn() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		config.User, config.Password, config.Host, config.Port, config.Database)
}

// Validate verifica se a configuração do MongoDB está completa
func (config MongoDBClientConfig) Validate() error {
	if config.User == "" || config.Password == "" || config.Host == "" || config.Port == "" || config.Database == "" {
		return fmt.Errorf("incomplete MongoDB configuration")
	}
	return nil
}
