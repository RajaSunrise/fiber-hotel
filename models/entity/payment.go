package entity

import "gorm.io/gorm"

// models payment
type Payment struct {
    gorm.Model
    BookingID      uint
    Amount         float64
    PaymentMethod  string
    TransactionID  string
}
