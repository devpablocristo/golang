package ginpkg

import "fmt"

type ginConfig struct {
	routerPort string
}

// NewGinConfig crea una nueva configuración para Gin con el puerto proporcionado.
func NewGinConfig(routerPort string) *ginConfig {
	return &ginConfig{
		routerPort: routerPort,
	}
}

// GetRouterPort devuelve el puerto del enrutador configurado.
func (config *ginConfig) GetRouterPort() string {
	return config.routerPort
}

// SetRouterPort establece el puerto del enrutador.
func (config *ginConfig) SetRouterPort(routerPort string) {
	config.routerPort = routerPort
}

// Validate valida la configuración de Gin.
func (config *ginConfig) Validate() error {
	if config.routerPort == "" {
		return fmt.Errorf("router port is not configured")
	}
	return nil
}
