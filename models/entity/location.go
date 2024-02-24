package entity

import "gorm.io/gorm"

// models location
type Location struct {
    gorm.Model
    Name            string
    Address         string
    Latitude        float64
    Longitude       float64
}