package monitoring

import (
	"github.com/gin-contrib/pprof" // Importa gin-contrib/pprof para integrar pprof con Gin
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/devpablocristo/golang/sdk/pkg/gin-gonic/gin/portspkg"
)

func Routes(ginpkg portspkg.GinClient, handler *GinHandler) {
	r := ginpkg.GetRouter()

	pprof.Register(r) // Registra las rutas de pprof en el enrutador de Gin

	// Ruta de Salud
	r.GET("/health", handler.Health)
	r.GET("/ping", handler.Ping)

	// Prometheus
	r.GET("/metrics", ginpkg.WrapH(promhttp.Handler()))

	// TODO: Falta Kong

}
