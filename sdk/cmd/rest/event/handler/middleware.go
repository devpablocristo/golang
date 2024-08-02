package event

import (
	"time"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/logger"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Registrar la solicitud entrante
		startTime := time.Now()
		logger.Infof("Incoming request: %s %s", c.Request.Method, c.Request.URL)

		// Procesar la solicitud
		c.Next()

		// Registrar la respuesta saliente
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		statusCode := c.Writer.Status()
		logger.Infof("Response: %d, Latency: %v", statusCode, latency)
	}
}
