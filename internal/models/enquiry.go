package models

import "gorm.io/gorm"

type Enquiry struct {
    gorm.Model
    UserID     uint   `json:"user_id"`
    PropertyID uint   `json:"property_id"`
    Message    string `json:"message"`
}
