package monitoring

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler representa el manejador de rutas para usuarios
type Handler struct{}

// NewHandler crea un nuevo manejador de rutas para usuarios
func NewHandler() *Handler {
	return &Handler{}
}

// Health verifica el estado del servicio y la conexión a la base de datos
func (h *Handler) Health(c *gin.Context) {
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
func (h *Handler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
