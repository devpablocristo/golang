package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	sdkports "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
	"github.com/devpablocristo/golang/sdk/services/analytics/gateways/analytics/dto"
	"github.com/devpablocristo/golang/sdk/services/analytics/internal/analytics/ports"
	"github.com/devpablocristo/golang/sdk/services/shared"
)

type Handler struct {
	ucs       ports.UseCases
	ginServer sdkports.Server
}

func NewRestHandler(ucs ports.UseCases, gsr sdkports.Server) *Handler {
	return &Handler{
		ucs:       ucs,
		ginServer: gsr,
	}
}

func (h *Handler) Start(apiVersion string, secret string) error {
	h.Routes(apiVersion, secret)
	return h.ginServer.RunServer()
}

func (h *Handler) Routes(apiVersion string, secret string) {
	//r := h.ginServer.GetRouter()

}

func (h *Handler) FakeCreateReport(c *gin.Context) {
	dir := "data"            // Aquí colocas la ruta de tu directorio
	filename := "event.json" // Nombre del archivo

	data, err := shared.LoadTestData(dir, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var metricsList dto.EventMetrics
	if err := json.Unmarshal(data, &metricsList); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = h.ucs.CreateReport(c.Request.Context(), metricsList.ToDomain()); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Report successfully created"})
}
