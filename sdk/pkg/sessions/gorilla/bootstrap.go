package sdksession

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang/sdk/pkg/sessions/gorilla/ports"
)

// Bootstrap inicializa el gestor de sesiones con la configuración necesaria
func Bootstrap() (ports.SessionManager, error) {
	config := newConfig(
		viper.GetString("GORILLA_SESSION_SECRET_KEY"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newSessionManager(config)
}
