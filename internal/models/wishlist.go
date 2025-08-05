package models

import "time"

type Wishlist struct {
    ID         uint      `json:"id" gorm:"primaryKey"`
    UserID     uint      `json:"user_id"`
    PropertyID uint      `json:"property_id"`
    CreatedAt  time.Time `json:"created_at"`
}
