package monitoring

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"

	wire "github.com/devpablocristo/golang/sdk/cmd/rest"
	basesetup "github.com/devpablocristo/golang/sdk/pkg/base-setup"
	gingonic "github.com/devpablocristo/golang/sdk/pkg/gin-gonic/gin"
	gmw "github.com/devpablocristo/golang/sdk/pkg/go-micro-web/v4"
)

func Routes(gingonic gingonic.GinClientPort, ms gmw.GoMicroClientPort) {
	r := gingonic.GetRouter()

	handler, err := wire.InitializeMonitoring()
	if err != nil {
		basesetup.MicroLogError("userHandler error: %v", err)
	}

	// Ruta de Salud
	r.GET("/health", handler.Health)
	r.GET("/ping", handler.Ping)

	// TODO: Probar prometheus
	r.GET("/metrics", gingonic.WrapH(promhttp.Handler()))

	// Integrar Go Micro y Gin
	// ms.GetService().Handle("/", r)
}
