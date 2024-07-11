package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/qh/events/internal/core"
)

type UserHandler struct {
	userUseCase core.UserUseCasePort
}

func NewUserHandler(userUseCase core.UserUseCasePort) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userUseCase.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Health(c *gin.Context) {
	// TODO implemntar
	// dbErr := h.ucs.CheckDatabaseConnection()
	// if dbErr != nil {
	//     c.JSON(http.StatusServiceUnavailable, gin.H{
	//         "status": "DOWN",
	//         "database": "unreachable",
	//     })
	//     return
	// }
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

func (h *UserHandler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
