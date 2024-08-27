package monitoring

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	ports "github.com/devpablocristo/golang/sdk/internal/core/monitoring/ports"
	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
)

type GinHandler struct {
	ucs       ports.UserUseCases
	ginServer sdkgin.Server
}

func NewGinHandler(u ports.UserUseCases, ginServer sdkgin.Server) *GinHandler {
	return &GinHandler{
		ucs:       u,
		ginServer: ginServer,
	}
}

func (h *GinHandler) Start(apiVersion string) error {
	h.Routes(apiVersion)
	return h.ginServer.RunServer()
}

func (h *GinHandler) Routes(apiVersion string) {
	r := h.ginServer.GetRouter()

	pprof.Register(r) // Registra las rutas de pprof en el enrutador de Gin

	// Ruta de Salud
	r.GET("/health", h.Health)
	r.GET("/ping", h.Ping)

	// Prometheus
	r.GET("/metrics", h.ginServer.WrapH(promhttp.Handler()))

	// TODO: Falta Kong
}

// Health verifica el estado del servicio y la conexión a la base de datos
func (h *GinHandler) Health(c *gin.Context) {
	// TODO: Implementar la verificación de la conexión a la base de datos
	// dbErr := h.ucs.CheckDatabaseConnection()
	// if dbErr != nil {
	//     c.JSON(http.StatusServiceUnavailable, gin.H{
	//         "status": "DOWN",
	//         "database": "unreachable",
	//     })
	//     return
	// }

	c.JSON(http.StatusOK, gin.H{
		"status": "up",
	})
}

// Ping responde con un mensaje "pong"
func (h *GinHandler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
