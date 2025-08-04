package repository

import (
    "github.com/Anujc07/real-estate/internal/models"
    "gorm.io/gorm"
)

type EnquiryRepository interface {
    Create(enquiry *models.Enquiry) error
    GetByPropertyID(propertyID uint) ([]models.Enquiry, error)
}

type enquiryRepository struct {
    db *gorm.DB
}

func NewEnquiryRepository(db *gorm.DB) EnquiryRepository {
    return &enquiryRepository{db}
}

func (r *enquiryRepository) Create(enquiry *models.Enquiry) error {
    return r.db.Create(enquiry).Error
}

func (r *enquiryRepository) GetByPropertyID(propertyID uint) ([]models.Enquiry, error) {
    var enquiries []models.Enquiry
    err := r.db.Where("property_id = ?", propertyID).Find(&enquiries).Error
    return enquiries, err
}
