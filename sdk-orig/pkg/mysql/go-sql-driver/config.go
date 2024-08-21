package gosqldriver

import "fmt"

type MySQLClientConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func (config MySQLClientConfig) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Database)
}

func (config MySQLClientConfig) Validate() error {
	if config.User == "" || config.Password == "" || config.Host == "" || config.Port == "" || config.Database == "" {
		return fmt.Errorf("incomplete MySQL configuration")
	}
	return nil
}