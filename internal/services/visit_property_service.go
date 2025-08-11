package service

import (
	"time"

	"github.com/Anujc07/real-estate/internal/models"
	"github.com/Anujc07/real-estate/internal/repository"
)

// VisitPropertyService interface to define the methods
type VisitPropertyService interface {
	Create(visit *models.VisitProperty) error
}

// visitpropertyService is the implementation of VisitPropertyService
type visitpropertyService struct {
	repo repository.VisitPropertyRepository
}

// NewVisitPropertyService returns an instance of VisitPropertyService
func NewVisitPropertyService(repo repository.VisitPropertyRepository) VisitPropertyService {
	return &visitpropertyService{
		repo: repo,
	}
}

// Create checks if user visited any property in last 7 days
// If yes, appends to existing entry, else creates a new one
func (s *visitpropertyService) Create(visit *models.VisitProperty) error {
	visit.VisitDate = time.Now()
	return s.repo.Create(visit)
}
