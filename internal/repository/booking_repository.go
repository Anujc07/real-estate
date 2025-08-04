package repository

import (
    "github.com/Anujc07/real-estate/internal/models"
    "gorm.io/gorm"
)

type BookingRepository interface {
    Create(booking *models.Booking) error
    FindAll() ([]models.Booking, error)
    FindByUser(userID uint) ([]models.Booking, error)
}

type bookingRepo struct {
    db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
    return &bookingRepo{db}
}

func (r *bookingRepo) Create(booking *models.Booking) error {
    return r.db.Create(booking).Error
}

func (r *bookingRepo) FindAll() ([]models.Booking, error) {
    var bookings []models.Booking
    err := r.db.Find(&bookings).Error
    return bookings, err
}

func (r *bookingRepo) FindByUser(userID uint) ([]models.Booking, error) {
    var bookings []models.Booking
    err := r.db.Where("user_id = ?", userID).Find(&bookings).Error
    return bookings, err
}
