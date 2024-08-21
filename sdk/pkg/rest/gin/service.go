package ginpkg

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	ports "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
)

var (
	instance  ports.Service
	once      sync.Once
	initError error
)

type service struct {
	router *gin.Engine
	config ports.Config
}

func NewService(config ports.Config) (ports.Service, error) {
	once.Do(func() {
		err := config.Validate()
		if err != nil {
			initError = err
			return
		}

		r := gin.Default()
		client := &service{
			config: config,
			router: r,
		}
		instance = client
	})
	return instance, initError
}

func GetInstance() (ports.Service, error) {
	if instance == nil {
		return nil, fmt.Errorf("gin client is not initialized")
	}
	return instance, nil
}

func (client *service) RunServer() error {
	return client.router.Run(":" + client.config.GetRouterPort())
}

func (client *service) GetRouter() *gin.Engine {
	return client.router
}

// WrapH envuelve un http.Handler en un gin.HandlerFunc.
func (client *service) WrapH(h http.Handler) gin.HandlerFunc {
	return gin.WrapH(h)
}
