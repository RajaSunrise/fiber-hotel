package entity

import (
	"time"
	"gorm.io/gorm"
)

// models user
type User struct {
    gorm.Model
    Name            string
    Username        string `gorm:"unique"`
    Email           string `gorm:"unique"`
    Password        string
    Address         string
    PhoneNumber     string
    MemberStatus    string
    BookingHistory  []Booking
    Reviews         []Review
}

// models Review
type Review struct {
    gorm.Model
    UserID      uint
    User        User
    RoomID      uint
    Room        Room
    Rating      int
    Comment     string
}

// models booking
type Booking struct {
    gorm.Model
    UserID         uint
    User           User
    RoomID         uint
    Room           Room
    CheckIn        time.Time
    CheckOut       time.Time
    TotalGuests    int
    PaymentStatus  string
    Payments       []Payment
}