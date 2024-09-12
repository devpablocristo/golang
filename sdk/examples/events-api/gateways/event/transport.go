package event

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/golang/sdk/examples/events-api/gateways/event/dto"
)

func decodeCreateEventRequest(c *gin.Context) (*dto.EventRequest, error) {
	var request dto.EventRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

func encodeResponse(c *gin.Context, response any) {
	c.JSON(http.StatusOK, response)
}

func errorHandler(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
