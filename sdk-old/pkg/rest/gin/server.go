package sdkgin

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	defs "github.com/devpablocristo/golang/sdk/pkg/rest/gin/defs"
)

var (
	instance  defs.Server
	once      sync.Once
	initError error
)

type server struct {
	router *gin.Engine
	config defs.Config
}

func newServer(config defs.Config) (defs.Server, error) {
	once.Do(func() {
		err := config.Validate()
		if err != nil {
			initError = err
			return
		}

		r := gin.Default()
		instance = &server{
			config: config,
			router: r,
		}
	})
	return instance, initError
}

func (server *server) RunServer() error {
	return server.router.Run(":" + server.config.GetRouterPort())
}

func (server *server) GetRouter() *gin.Engine {
	return server.router
}

func (server *server) GetApiVersion() string {
	return server.config.GetApiVersion()
}

// WrapH envuelve un http.Handler en un gin.HandlerFunc.
func (server *server) WrapH(h http.Handler) gin.HandlerFunc {
	return gin.WrapH(h)
}
