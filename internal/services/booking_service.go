package service

import (
    "github.com/Anujc07/real-estate/internal/models"
    "github.com/Anujc07/real-estate/internal/repository"
)

type BookingService interface {
    Book(booking *models.Booking) error
    GetAllBookings() ([]models.Booking, error)
    GetBookingsByUser(userID uint) ([]models.Booking, error)
}

type bookingService struct {
    repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) BookingService {
    return &bookingService{repo}
}

func (s *bookingService) Book(booking *models.Booking) error {
    // You could add business logic here
    booking.Status = "Pending"
    return s.repo.Create(booking)
}

func (s *bookingService) GetAllBookings() ([]models.Booking, error) {
    return s.repo.FindAll()
}

func (s *bookingService) GetBookingsByUser(userID uint) ([]models.Booking, error) {
    return s.repo.FindByUser(userID)
}
