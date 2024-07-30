package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	pst "github.com/devpablocristo/dive-challenge/cmd/api/handlers/presenter"
	ucs "github.com/devpablocristo/dive-challenge/internal/core"
)

type RestHandler struct {
	ucs ucs.UseCasePort
}

func NewRestHandler(ucs ucs.UseCasePort) *RestHandler {
	return &RestHandler{ucs: ucs}
}

func (h *RestHandler) GetLTP(c *gin.Context) {
	pairs := c.QueryArray("pair")

	if len(pairs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'pair' query parameter"})
		return
	}

	ltp, err := h.ucs.GetLTP(c.Request.Context(), pairs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := pst.ToResponseLTPList(ltp)
	c.JSON(http.StatusOK, response)
}
