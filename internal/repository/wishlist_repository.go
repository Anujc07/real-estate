package repository

import (
    "github.com/Anujc07/real-estate/internal/models"
    "gorm.io/gorm"
)

type WishlistRepository interface {
    Add(item *models.Wishlist) error
    Remove(userID, propertyID uint) error
    GetByUser(userID uint) ([]models.Wishlist, error)
}

type wishlistRepository struct {
    db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) WishlistRepository {
    return &wishlistRepository{db}
}

func (r *wishlistRepository) Add(item *models.Wishlist) error {
    return r.db.Create(item).Error
}

func (r *wishlistRepository) Remove(userID, propertyID uint) error {
    return r.db.Where("user_id = ? AND property_id = ?", userID, propertyID).Delete(&models.Wishlist{}).Error
}

func (r *wishlistRepository) GetByUser(userID uint) ([]models.Wishlist, error) {
    var items []models.Wishlist
    err := r.db.Where("user_id = ?", userID).Find(&items).Error
    return items, err
}
