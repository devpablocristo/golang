package ginsetup

import (
	"fmt"

	"github.com/spf13/viper"

	ginpkg "github.com/devpablocristo/golang/sdk/pkg/gin-gonic/gin"
	"github.com/devpablocristo/golang/sdk/pkg/gin-gonic/gin/portspkg"
)

func NewGinInstance() (portspkg.GinClient, error) {
	// Crear la configuración utilizando viper para leer las variables de entorno
	config := ginpkg.NewGinConfig(viper.GetString("ROUTER_PORT"))
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create gin config: %w", err)
	// }

	// Inicializar el cliente Gin con la configuración
	if err := ginpkg.InitializeGinClient(config); err != nil {
		return nil, fmt.Errorf("failed to initialize gin client: %w", err)
	}

	// Devolver la instancia del cliente Gin
	return ginpkg.GetGinInstance()
}
