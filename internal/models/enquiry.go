package models

import "gorm.io/gorm"

type Enquiry struct {
    gorm.Model
    UserID     uint   `json:"user_id"`
    PropertyID uint   `json:"property_id"`
    Message       string `json:"message"`
    InterestedCity string `json:"interested_city"`
    InterestedArea string `json:"interested_area"`
}
