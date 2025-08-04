package service

import (
    "github.com/Anujc07/real-estate/internal/models"
    "github.com/Anujc07/real-estate/internal/repository"
)

type ReviewService interface {
    CreateReview(review *models.Review) error
    GetReviews(propertyID uint) ([]models.Review, error)
}

type reviewService struct {
    repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) ReviewService {
    return &reviewService{repo}
}

func (s *reviewService) CreateReview(review *models.Review) error {
    return s.repo.Create(review)
}

func (s *reviewService) GetReviews(propertyID uint) ([]models.Review, error) {
    return s.repo.GetByProperty(propertyID)
}
