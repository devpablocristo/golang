package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/golang-sdk/internal/core"
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

// type RestHandler struct {
// 	ucs ucs.UseCasePort
// }

// func NewRestHandler(ucs ucs.UseCasePort) *RestHandler {
// 	return &RestHandler{ucs: ucs}
// }

// func (h *RestHandler) HelloWorld(c *gin.Context) {
// 	str := "Hello, World!!! Olá Mundo!!! Hola Mundo!!!"
// 	c.JSON(http.StatusOK, str)
// }

// func (h *RestHandler) GetUser(c *gin.Context) {
// 	id := c.Param("id")
// 	ctx := c.Request.Context()

// 	user, err := h.ucs.GetUser(ctx, id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// func (h *RestHandler) DeleteUser(c *gin.Context) {
// 	id := c.Param("id")
// 	ctx := c.Request.Context()

// 	err := h.ucs.DeleteUser(ctx, id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
// 		return
// 	}
// 	c.Status(http.StatusOK)
// }

// func (h *RestHandler) ListUsers(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	users, err := h.ucs.ListUsers(ctx)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, users)
// }

// func (h *RestHandler) UpdateUser(c *gin.Context) {
// 	var user usr.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	id := c.Param("id")
// 	ctx := c.Request.Context()

// 	err := h.ucs.UpdateUser(ctx, &user, id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating user"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// func (h *RestHandler) CreateUser(c *gin.Context) {
// 	var user usr.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx := c.Request.Context()
// 	err := h.ucs.CreateUser(ctx, &user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating user"})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, user)
// }
