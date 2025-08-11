package models

import (
	"time"
)

type VisitProperty struct {
	ID uint `gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	PropertyIDs []uint `gorm:"type:json" json:"property_ids"`
	VisitDate time.Time `json:"visit_date"`
}