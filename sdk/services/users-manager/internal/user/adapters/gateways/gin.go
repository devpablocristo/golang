package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	mware "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
	entities "github.com/devpablocristo/golang/sdk/services/users-manager/internal/user/core/entities"
	ports "github.com/devpablocristo/golang/sdk/services/users-manager/internal/user/core/ports"
)

type GinHandler struct {
	ucs       ports.UseCases
	ginServer sdkgin.Server
}

func NewGinHandler(u ports.UseCases, ginServer sdkgin.Server) *GinHandler {
	return &GinHandler{
		ucs:       u,
		ginServer: ginServer,
	}
}
func (h *GinHandler) Start(apiVersion string, secret string) error {
	h.Routes(apiVersion, secret)
	return h.ginServer.RunServer()
}

func (h *GinHandler) Routes(apiVersion string, secret string) {
	router := h.ginServer.GetRouter()

	apiPrefix := "/api/" + apiVersion

	router.GET(apiPrefix+"/health", h.Health)

	authorized := router.Group(apiPrefix + "/user/protected")
	authorized.Use(mware.ValidateJwt(secret))
	{
		authorized.GET("/user-protected", h.CreateUser)
	}
}

func (h *GinHandler) CreateUser(c *gin.Context) {
	var user entities.User
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
	uuid := c.Param("uuid")
	user, err := h.ucs.GetUserByUUID(c.Request.Context(), uuid)
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
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	uuid := c.Param("uuid")
	if err := h.ucs.UpdateUser(c.Request.Context(), &user, uuid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *GinHandler) DeleteUser(c *gin.Context) {
	uuid := c.Param("uuid")
	err := h.ucs.DeleteUser(c.Request.Context(), uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *GinHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
