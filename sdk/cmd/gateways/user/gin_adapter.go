package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/golang/sdk/internal/core"
	"github.com/devpablocristo/golang/sdk/internal/core/user"
)

type GinHandler struct {
	ucs core.UserUseCasesPort
}

func NewGinHandler(ucs core.UserUseCasesPort) *GinHandler {
	return &GinHandler{
		ucs: ucs,
	}
}

func (h *GinHandler) CreateUser(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.ucs.CreateUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *GinHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.ucs.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *GinHandler) ListUsers(c *gin.Context) {
	users, err := h.ucs.ListUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *GinHandler) UpdateUser(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	id := c.Param("id")
	if err := h.ucs.UpdateUser(c.Request.Context(), &user, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *GinHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := h.ucs.DeleteUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
