package repository

import (
	"github.com/Anujc07/real-estate/internal/models"
	"gorm.io/gorm"
	"time"
)

type VisitPropertyRepository interface {
	Create(visit *models.VisitProperty) error
}

type visitpropertyRepo struct {
	DB *gorm.DB
}

func NewVisitPropertyRepository(db *gorm.DB) *visitpropertyRepo {
	return &visitpropertyRepo{
		DB: db,
	}
}

func (r *visitpropertyRepo) Create(visit *models.VisitProperty) error {
	var existingVisit models.VisitProperty
	oneWeekAgo := time.Now().AddDate(0, 0, -7)

	err := r.DB.Where("user_id = ? AND visit_date >= ?", visit.UserID, oneWeekAgo).First(&existingVisit).Error

	if err == nil {
		existingVisit.PropertyIDs = append(existingVisit.PropertyIDs, visit.PropertyIDs...)
		return r.DB.Save(&existingVisit).Error
	}
	return r.DB.Create(visit).Error
}