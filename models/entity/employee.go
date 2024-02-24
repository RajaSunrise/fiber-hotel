package entity

import "gorm.io/gorm"

// models employee
type Employee struct {
    gorm.Model
    Name            string
    Position        string
}