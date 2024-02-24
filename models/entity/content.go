package entity

import "gorm.io/gorm"

// models content
type Content struct {
    gorm.Model
    Title           string
    Body            string `gorm:"type:text"`
}