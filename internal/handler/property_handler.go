package handler

import (
    "net/http"
    "strconv"

    "github.com/Anujc07/real-estate/internal/models"
    "github.com/Anujc07/real-estate/internal/services"
    "github.com/gin-gonic/gin"
)

type PropertyHandler struct {
    service service.PropertyService
}

func NewPropertyHandler(service service.PropertyService) *PropertyHandler {
    return &PropertyHandler{service}
}

func (h *PropertyHandler) Add(c *gin.Context) {
    var prop models.Property
    if err := c.ShouldBindJSON(&prop); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.service.Add(&prop); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add property"})
        return
    }
    c.JSON(http.StatusCreated, prop)
}

func (h *PropertyHandler) GetAll(c *gin.Context) {
    props, err := h.service.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch properties"})
        return
    }
    c.JSON(http.StatusOK, props)
}

func (h *PropertyHandler) GetByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid property ID"})
        return
    }
    prop, err := h.service.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
        return
    }
    c.JSON(http.StatusOK, prop)
}
