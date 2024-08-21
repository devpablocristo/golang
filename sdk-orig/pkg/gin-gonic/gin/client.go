package ginpkg

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/devpablocristo/golang/sdk/pkg/gin-gonic/gin/portspkg"
	"github.com/gin-gonic/gin"
)

var (
	instance portspkg.GinClient
	once     sync.Once
	errInit  error
)

type ginClient struct {
	router *gin.Engine
	config *ginConfig
}

// InitializeGinClient inicializa el cliente Gin como un singleton.
func InitializeGinClient(config *ginConfig) error {
	once.Do(func() {
		err := config.Validate()
		if err != nil {
			errInit = err
			return
		}

		r := gin.Default()
		client := &ginClient{
			config: config,
			router: r,
		}
		instance = client
	})
	return errInit
}

// GetGinInstance devuelve la instancia del cliente Gin.
func GetGinInstance() (portspkg.GinClient, error) {
	if instance == nil {
		return nil, fmt.Errorf("gin client is not initialized")
	}
	return instance, nil
}

// RunServer inicia el servidor Gin en el puerto configurado.
func (client *ginClient) RunServer() error {
	return client.router.Run(":" + client.config.GetRouterPort())
}

// GetRouter devuelve el enrutador Gin.
func (client *ginClient) GetRouter() *gin.Engine {
	return client.router
}

// WrapH envuelve un http.Handler en un gin.HandlerFunc.
func (client *ginClient) WrapH(h http.Handler) gin.HandlerFunc {
	return gin.WrapH(h)
}
