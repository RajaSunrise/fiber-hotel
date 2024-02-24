package entity

import (
	"time"
	"gorm.io/gorm"
)

// models promotion
type Promotion struct {
    gorm.Model
    Code            string `gorm:"unique"`
    Description     string
    ValidFrom       time.Time
    ValidTo         time.Time
}
