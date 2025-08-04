package repository

import (
    "github.com/Anujc07/real-estate/internal/models"
    "gorm.io/gorm"
)

type ReviewRepository interface {
    Create(review *models.Review) error
    GetByProperty(propertyID uint) ([]models.Review, error)
}

type reviewRepository struct {
    db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
    return &reviewRepository{db}
}

func (r *reviewRepository) Create(review *models.Review) error {
    return r.db.Create(review).Error
}

func (r *reviewRepository) GetByProperty(propertyID uint) ([]models.Review, error) {
    var reviews []models.Review
    err := r.db.Where("property_id = ?", propertyID).Find(&reviews).Error
    return reviews, err
}
