package repository

import (
    "github.com/Anujc07/real-estate/internal/models"
    "gorm.io/gorm"
)

type PropertyRepository interface {
    Create(property *models.Property) error
    FindAll() ([]models.Property, error)
    FindByID(id uint) (models.Property, error)
}

type propertyRepository struct {
    db *gorm.DB
}

func NewPropertyRepository(db *gorm.DB) PropertyRepository {
    return &propertyRepository{db}
}

func (r *propertyRepository) Create(property *models.Property) error {
    return r.db.Create(property).Error
}

func (r *propertyRepository) FindAll() ([]models.Property, error) {
    var props []models.Property
    err := r.db.Find(&props).Error
    return props, err
}

func (r *propertyRepository) FindByID(id uint) (models.Property, error) {
    var prop models.Property
    err := r.db.First(&prop, id).Error
    return prop, err
}
