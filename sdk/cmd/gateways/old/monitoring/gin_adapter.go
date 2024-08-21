package monitoring

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GinHandler representa el manejador de rutas para usuarios
type GinHandler struct{}

// NewGinHandler crea un nuevo manejador de rutas para usuarios
func NewGinHandler() *GinHandler {
	return &GinHandler{}
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
