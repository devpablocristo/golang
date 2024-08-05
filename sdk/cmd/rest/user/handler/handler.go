package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/golang/sdk/internal/core"
	"github.com/devpablocristo/golang/sdk/internal/core/user"
)

// Handler representa el manejador de rutas para usuarios
type Handler struct {
	ucs core.UserUseCasePort
}

// NewHandler crea un nuevo manejador de rutas para usuarios
func NewHandler(ucs core.UserUseCasePort) *Handler {
	return &Handler{
		ucs: ucs,
	}
}

// GetUser obtiene un usuario por su ID
func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.ucs.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser elimina un usuario por su ID
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	err := h.ucs.DeleteUser(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.Status(http.StatusNoContent) // Usar 204 No Content para eliminaciones exitosas
}

// ListUsers lista todos los usuarios
func (h *Handler) ListUsers(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := h.ucs.ListUsers(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// UpdateUser actualiza la información de un usuario
func (h *Handler) UpdateUser(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	id := c.Param("id")
	ctx := c.Request.Context()
	if err := h.ucs.UpdateUser(ctx, &user, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser crea un nuevo usuario
func (h *Handler) CreateUser(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ctx := c.Request.Context()
	if err := h.ucs.CreateUser(ctx, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}
