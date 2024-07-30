package handler

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	tst "github.com/devpablocristo/qh/events/cmd/api/tests"
	ucs "github.com/devpablocristo/qh/events/internal/core"
)

type RestHandler struct {
	ucs ucs.UseCasePort
}

func NewRestHandler(ucs ucs.UseCasePort) *RestHandler {
	return &RestHandler{
		ucs: ucs,
	}
}

func (h *RestHandler) FakeCreateEvent(c *gin.Context) {
	data, err := tst.LoadTestData()
	if err != nil {
		errorHandler(c, err)
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(data))

	dto, err := decodeCreateEventRequest(c)
	if err != nil {
		errorHandler(c, err)
		return
	}

	if err = h.ucs.CreateEvent(c.Request.Context(), dto.ToDomain()); err != nil {
		errorHandler(c, err)
		return
	}

	response := CreateEventResponse{
		Message: "Event created successfully",
	}
	encodeResponse(c, response)
}

func (h *RestHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}
