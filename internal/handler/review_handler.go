package handler

import (
    "net/http"
    "strconv"

    "github.com/Anujc07/real-estate/internal/models"
    "github.com/Anujc07/real-estate/internal/services"
    "github.com/gin-gonic/gin"
)

type ReviewHandler struct {
    service service.ReviewService
}

func NewReviewHandler(service service.ReviewService) *ReviewHandler {
    return &ReviewHandler{service}
}

// POST /reviews
func (h *ReviewHandler) Create(c *gin.Context) {
    var review models.Review
    if err := c.ShouldBindJSON(&review); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if review.Rating < 1 || review.Rating > 5 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 1 and 5"})
        return
    }

    if err := h.service.CreateReview(&review); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit review"})
        return
    }

    c.JSON(http.StatusCreated, review)
}

// GET /reviews/property/:id
func (h *ReviewHandler) GetByProperty(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid property ID"})
        return
    }

    reviews, err := h.service.GetReviews(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get reviews"})
        return
    }

    c.JSON(http.StatusOK, reviews)
}
