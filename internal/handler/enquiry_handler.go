package handler

import (
    "net/http"
    "strconv"

    "github.com/Anujc07/real-estate/internal/models"
    "github.com/Anujc07/real-estate/internal/services"
    "github.com/gin-gonic/gin"
)

type EnquiryHandler struct {
    service service.EnquiryService
}

func NewEnquiryHandler(service service.EnquiryService) *EnquiryHandler {
    return &EnquiryHandler{service}
}

// POST /enquiries
func (h *EnquiryHandler) Create(c *gin.Context) {
    var enquiry models.Enquiry
    if err := c.ShouldBindJSON(&enquiry); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.CreateEnquiry(&enquiry); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send enquiry"})
        return
    }

    c.JSON(http.StatusCreated, enquiry)
}

// GET /enquiries/property/:id
func (h *EnquiryHandler) GetByPropertyID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid property ID"})
        return
    }

    enquiries, err := h.service.GetEnquiries(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get enquiries"})
        return
    }

    c.JSON(http.StatusOK, enquiries)
}
