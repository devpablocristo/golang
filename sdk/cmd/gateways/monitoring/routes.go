package monitoring

import (
	"github.com/gin-contrib/pprof" // Importa gin-contrib/pprof para integrar pprof con Gin
	"github.com/prometheus/client_golang/prometheus/promhttp"

	gingonic "github.com/devpablocristo/golang/sdk/pkg/gin-gonic/gin"
)

func Routes(gingonic gingonic.GinClientPort, handler *Handler) {
	r := gingonic.GetRouter()

	pprof.Register(r) // Registra las rutas de pprof en el enrutador de Gin

	// Ruta de Salud
	r.GET("/health", handler.Health)
	r.GET("/ping", handler.Ping)

	// Prometheus
	r.GET("/metrics", gingonic.WrapH(promhttp.Handler()))

	// TODO: Falta Kong

}
