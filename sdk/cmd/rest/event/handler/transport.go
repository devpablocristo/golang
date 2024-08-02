package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func decodeCreateEventRequest(c *gin.Context) (*EventRequest, error) {
	var request EventRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

func encodeResponse(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}

func errorHandler(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
