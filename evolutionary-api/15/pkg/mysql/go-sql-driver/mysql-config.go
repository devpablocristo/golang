package gosqldriver

import (
	"fmt"
)

// MySQLClientConfig contiene la configuración necesaria para conectarse a una base de datos MySQL
type MySQLClientConfig struct {
	User     string // Usuario de la base de datos
	Password string // Contraseña del usuario
	Host     string // Host donde se encuentra la base de datos
	Port     string // Puerto en el que escucha la base de datos
	Database string // Nombre de la base de datos
}

// dsn genera el Data Source Name (DSN) a partir de la configuración proporcionada
func (config MySQLClientConfig) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Database)
}
