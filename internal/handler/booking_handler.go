package handler

import (
    "net/http"
    "strconv"

    "github.com/Anujc07/real-estate/internal/models"
    "github.com/Anujc07/real-estate/internal/services"
    "github.com/gin-gonic/gin"
)

type BookingHandler struct {
    service service.BookingService
}

func NewBookingHandler(service service.BookingService) *BookingHandler {
    return &BookingHandler{service}
}

func (h *BookingHandler) BookProperty(c *gin.Context) {
    var booking models.Booking
    if err := c.ShouldBindJSON(&booking); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.Book(&booking); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Booking failed"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Booking created", "booking": booking})
}

func (h *BookingHandler) GetAll(c *gin.Context) {
    bookings, err := h.service.GetAllBookings()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
        return
    }

    c.JSON(http.StatusOK, bookings)
}

func (h *BookingHandler) GetByUser(c *gin.Context) {
    userIDStr := c.Param("user_id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    bookings, err := h.service.GetBookingsByUser(uint(userID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
        return
    }

    c.JSON(http.StatusOK, bookings)
}
