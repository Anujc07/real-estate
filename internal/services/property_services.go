package service

import (
    "github.com/Anujc07/real-estate/internal/models"
    "github.com/Anujc07/real-estate/internal/repository"
)

type PropertyService interface {
    Add(property *models.Property) error
    GetAll() ([]models.Property, error)
    GetByID(id uint) (models.Property, error)
}

type propertyService struct {
    repo repository.PropertyRepository
}

func NewPropertyService(repo repository.PropertyRepository) PropertyService {
    return &propertyService{repo}
}

func (s *propertyService) Add(property *models.Property) error {
    return s.repo.Create(property)
}

func (s *propertyService) GetAll() ([]models.Property, error) {
    return s.repo.FindAll()
}

func (s *propertyService) GetByID(id uint) (models.Property, error) {
    return s.repo.FindByID(id)
}
