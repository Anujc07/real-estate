package models

import "time"

type Property struct {
    ID          uint      `gorm:"primaryKey"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Price       float64   `json:"price"`
    Location    string    `json:"location"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    AreaSqft    int     `json:"area_sqft"`
    Bedrooms    int        `json:"bedrooms"`
    Bathrooms   int        `json:"bathrooms"`
    PropertyType string     `json:"property_type"`
    ListedBy     string     `json:"listed_by"`
    ImagesUrl    []string   `gorm:"type:json" json:"images_urls"`
}
