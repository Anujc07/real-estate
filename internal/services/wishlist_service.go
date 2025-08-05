package service

import (
    "github.com/Anujc07/real-estate/internal/models"
    "github.com/Anujc07/real-estate/internal/repository"
)

type WishlistService interface {
    AddToWishlist(item *models.Wishlist) error
    RemoveFromWishlist(userID, propertyID uint) error
    GetWishlist(userID uint) ([]models.Wishlist, error)
}

type wishlistService struct {
    repo repository.WishlistRepository
}

func NewWishlistService(repo repository.WishlistRepository) WishlistService {
    return &wishlistService{repo}
}

func (s *wishlistService) AddToWishlist(item *models.Wishlist) error {
    return s.repo.Add(item)
}

func (s *wishlistService) RemoveFromWishlist(userID, propertyID uint) error {
    return s.repo.Remove(userID, propertyID)
}

func (s *wishlistService) GetWishlist(userID uint) ([]models.Wishlist, error) {
    return s.repo.GetByUser(userID)
}
