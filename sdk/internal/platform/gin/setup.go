package ginsetup

import (
	"fmt"

	"github.com/spf13/viper"

	gin "github.com/devpablocristo/qh/events/pkg/gin-gonic/gin"
)

func NewGinInstance() (gin.GinClientPort, error) {
	config := gin.GinConfig{
		RouterPort: viper.GetString("ROUTER_PORT"),
	}

	if config.RouterPort == "" {
		return nil, fmt.Errorf("router port is not configured")
	}

	if err := gin.InitializeGinClient(config); err != nil {
		return nil, err
	}

	return gin.GetGinInstance()
}
