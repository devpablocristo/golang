package monitoring

import (
	"github.com/gin-contrib/pprof" // Importa gin-contrib/pprof para integrar pprof con Gin
	"github.com/prometheus/client_golang/prometheus/promhttp"

	wire "github.com/devpablocristo/golang/sdk/cmd/rest"
	basesetup "github.com/devpablocristo/golang/sdk/pkg/base-setup"
	gingonic "github.com/devpablocristo/golang/sdk/pkg/gin-gonic/gin"
)

func Routes(gingonic gingonic.GinClientPort) {
	r := gingonic.GetRouter()

	handler, err := wire.InitializeMonitoring()
	if err != nil {
		basesetup.MicroLogError("userHandler error: %v", err)
	}

	pprof.Register(r) // Registra las rutas de pprof en el enrutador de Gin

	// Ruta de Salud
	r.GET("/health", handler.Health)
	r.GET("/ping", handler.Ping)

	// Prometheus
	r.GET("/metrics", gingonic.WrapH(promhttp.Handler()))

	// TODO: Falta Kong

}
