package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	ttools "github.com/devpablocristo/golang/sdk/cmd/rest/testing-tools"
	ucs "github.com/devpablocristo/golang/sdk/internal/core"
)

type Handler struct {
	ucs ucs.UseCasesPort
}

func NewRestHandler(ucs ucs.UseCasesPort) *Handler {
	return &Handler{
		ucs: ucs,
	}
}

func (h *Handler) FakeCreateReport(c *gin.Context) {
	data, err := ttools.LoadTestData("test")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var metricsList EventMetricsDTO
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
