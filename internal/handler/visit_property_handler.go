package handler

import (
	"net/http"

	"github.com/Anujc07/real-estate/internal/models"
	"github.com/Anujc07/real-estate/internal/services"
	"github.com/gin-gonic/gin"
)

// VisitPropertyHandler struct
type VisitPropertyHandler struct {
	service service.VisitPropertyService
}

// NewVisitPropertyHandler returns a handler instance
func NewVisitPropertyHandler(service service.VisitPropertyService) *VisitPropertyHandler {
	return &VisitPropertyHandler{service: service}
}

// CreateVisit handles POST /visit
func (h *VisitPropertyHandler) CreateVisit(c *gin.Context) {
	var visit models.VisitProperty

	if err := c.ShouldBindJSON(&visit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&visit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store visit"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Visit recorded"})
}
