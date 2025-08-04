package service

import (
    "github.com/Anujc07/real-estate/internal/models"
    "github.com/Anujc07/real-estate/internal/repository"
)

type EnquiryService interface {
    CreateEnquiry(enquiry *models.Enquiry) error
    GetEnquiries(propertyID uint) ([]models.Enquiry, error)
}

type enquiryService struct {
    repo repository.EnquiryRepository
}

func NewEnquiryService(repo repository.EnquiryRepository) EnquiryService {
    return &enquiryService{repo}
}

func (s *enquiryService) CreateEnquiry(enquiry *models.Enquiry) error {
    return s.repo.Create(enquiry)
}

func (s *enquiryService) GetEnquiries(propertyID uint) ([]models.Enquiry, error) {
    return s.repo.GetByPropertyID(propertyID)
}
