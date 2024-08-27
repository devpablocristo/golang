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
	ucs ports.UseCases
	gs  sdkgin.Server
}

func NewGinHandler(u ports.UseCases, s sdkgin.Server) *GinHandler {
	return &GinHandler{
		ucs: u,
		gs:  s,
	}
}

func (h *GinHandler) Start(apiVersion string) error {
	h.Routes(apiVersion)
	return h.gs.RunServer()
}

func (h *GinHandler) Routes(apiVersion string) {
	r := h.gs.GetRouter()

	// Registra las rutas de pprof en el enrutador de Gin
	pprof.Register(r)

	apiPrefix := "/api/" + apiVersion

	// Rutas de Salud
	r.GET(apiPrefix+"/health", h.Health)
	r.GET(apiPrefix+"/db-health", h.DbHealth)
	r.GET(apiPrefix+"/ping", h.Ping)

	// Prometheus
	r.GET("/metrics", h.gs.WrapH(promhttp.Handler()))
}

// Health verifica el estado del servicio y la conexión a la base de datos
func (h *GinHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

func (h *GinHandler) DbHealth(c *gin.Context) {
	// Verificar la conexión a la base de datos
	dbErr := h.ucs.CheckDbConn(c.Request.Context())
	if dbErr != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":   "DOWN",
			"database": "unreachable",
		})
		return
	}

	// Si la base de datos está accesible
	c.JSON(http.StatusOK, gin.H{
		"status":   "UP",
		"database": "reachable",
	})
}

// Ping responde con un mensaje "pong"
func (h *GinHandler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
