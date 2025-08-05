package handler

import (
    "net/http"
    "strconv"

    "github.com/Anujc07/real-estate/internal/models"
    "github.com/Anujc07/real-estate/internal/services"
    "github.com/gin-gonic/gin"
)

type WishlistHandler struct {
    service service.WishlistService
}

func NewWishlistHandler(service service.WishlistService) *WishlistHandler {
    return &WishlistHandler{service}
}

// POST /wishlist
func (h *WishlistHandler) Add(c *gin.Context) {
    var item models.Wishlist
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.AddToWishlist(&item); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to wishlist"})
        return
    }

    c.JSON(http.StatusCreated, item)
}

// DELETE /wishlist/:user_id/:property_id
func (h *WishlistHandler) Remove(c *gin.Context) {
    userID, _ := strconv.Atoi(c.Param("user_id"))
    propertyID, _ := strconv.Atoi(c.Param("property_id"))

    if err := h.service.RemoveFromWishlist(uint(userID), uint(propertyID)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove from wishlist"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Removed successfully"})
}

// GET /wishlist/:user_id
func (h *WishlistHandler) Get(c *gin.Context) {
    userID, _ := strconv.Atoi(c.Param("user_id"))

    list, err := h.service.GetWishlist(uint(userID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch wishlist"})
        return
    }

    c.JSON(http.StatusOK, list)
}
