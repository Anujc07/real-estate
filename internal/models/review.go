package models

import "time"

type Review struct {
    ID         uint      `json:"id" gorm:"primaryKey"`
    PropertyID uint      `json:"property_id"`
    UserID     uint      `json:"user_id"`
    Rating     int       `json:"rating"`  // 1 to 5
    Comment    string    `json:"comment"`
    CreatedAt  time.Time `json:"created_at"`
}
