package entity

import "gorm.io/gorm"

type Room struct {
    gorm.Model
    RoomNumber      string
    Type            string
    Price           float64
    Facilities      []*Facility `gorm:"many2many:room_facilities;"`
    Bookings        []Booking
}

type Facility struct {
    gorm.Model
    Name            string
    Description     string
    Rooms           []*Room `gorm:"many2many:room_facilities;"`
}
