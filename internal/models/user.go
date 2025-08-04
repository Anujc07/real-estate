package models

import (
	"time"
	// "gorm.io/gorm"
)

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string
    Email    string `gorm:"unique"`
    Password string
    Role     string `gorm:"default:'customer'"` // New: admin, agent, customer
    CreatedAt time.Time
}

