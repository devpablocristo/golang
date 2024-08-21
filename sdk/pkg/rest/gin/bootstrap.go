package ginpkg

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
)

func Bootstrap() (ports.Service, error) {
	config := NewConfig(viper.GetString("ROUTER_PORT"))

	return NewService(config)
}
