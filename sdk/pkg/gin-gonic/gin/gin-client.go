package gingonic

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	instance GinClientPort
	once     sync.Once
	errInit  error
)

type GinClient struct {
	router *gin.Engine
	config GinConfig
}

func InitializeGinClient(config GinConfig) error {
	once.Do(func() {
		r := gin.Default()
		client := &GinClient{
			config: config,
			router: r,
		}
		instance = client
	})
	return errInit
}

func GetGinInstance() (GinClientPort, error) {
	if instance == nil {
		return nil, fmt.Errorf("gin client is not initialized")
	}
	return instance, nil
}

func (client *GinClient) RunServer() error {
	if client.config.RouterPort == "" {
		return fmt.Errorf("router port is not configured")
	}
	return client.router.Run(":" + client.config.RouterPort)
}

func (client *GinClient) GetRouter() *gin.Engine {
	return client.router
}

func (client *GinClient) WrapH(h http.Handler) gin.HandlerFunc {
	return gin.WrapH(h)
}
