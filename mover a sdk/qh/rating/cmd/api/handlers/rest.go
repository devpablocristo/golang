package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	ucs "github.com/devpablocristo/qh/rating/internal/core"
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

	response := ToResponseLTPList(ltp)
	c.JSON(http.StatusOK, response)
}

func (h *RestHandler) CreateRating(c *gin.Context) {
	var ratingDTO RatingDTO
	if err := c.ShouldBindJSON(&ratingDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usecase := rating.NewRatingUseCase()
	err := usecase.CreateRating(ToRating())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Rating created successfully"})
}

func (h *RestHandler) GetRating(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	usecase := rating.NewRatingUseCase()
	rate, err := usecase.GetRating(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rating not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rating": handler.FromDomain(rate)})
}
