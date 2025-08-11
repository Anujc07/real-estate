package models

import (
	"time"
)

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string
    Email    string `gorm:"unique"`
    Password string
    Role     string `gorm:"default:'customer'"` // New: admin, agent, customer
    CreatedAt time.Time
    UpdatedAt time.Time
    City      string `gorm:"column:city"`
    State     string `gorm:"column:state"`
    IsActive  bool   `gorm:"default:true"` // New: to track if user is active
}

