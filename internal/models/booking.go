package models

import "gorm.io/gorm"

type Booking struct {
    gorm.Model
    UserID     uint   `json:"user_id"`
    PropertyID uint   `json:"property_id"`
    Status     string `json:"status"` // Pending, Confirmed, Cancelled
}
