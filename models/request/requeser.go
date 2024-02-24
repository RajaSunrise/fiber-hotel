package request

import (
    "gorm.io/gorm"
    "time"
)


type User struct {
    gorm.Model
    Name           string `json:"name" validate:"required,min=2,max=100"`
    Username       string `json:"username" validate:"required,min=8,max=12"`
    Email          string `json:"email" validate:"required,email"`
    Password       string `json:"password" validate:"required,min=6,max=64"`
    Address        string `json:"address"`
    PhoneNumber    string `json:"phone_number"`
    MemberStatus   string `json:"member_status"`
    Reviews        []Review    `json:"reviews"`
    BookingHistory []Booking   `json:"booking_history"`
}

// Facility model
type Facility struct {
    gorm.Model
    Name        string    `json:"name" validate:"required"`
    Description string    `json:"description"`
    Rooms       []Room    `gorm:"many2many:facility_rooms;" json:"rooms"`
}

// Room model
type Room struct {
    gorm.Model
    RoomNumber string     `json:"room_number" validate:"required"`
    Type       string     `json:"type"`
    Price      float64    `json:"price" validate:"required,gt=0"`
    Facilities []Facility `gorm:"many2many:facility_rooms;" json:"facilities"`
    Bookings   []Booking  `json:"bookings"`
}


// Review model
type Review struct {
    gorm.Model
    UserID uint      `json:"user_id"`
    User   User      `json:"user"`
    RoomID uint      `json:"room_id"`
    Room   Room      `json:"room"`
    Rating int       `json:"rating" validate:"required,gte=1,lte=5"`
    Comment string   `json:"comment"`
}

// Booking model
type Booking struct {
    gorm.Model
    UserID        uint      `json:"user_id"`
    User          User      `json:"user"`
    RoomID        uint      `json:"room_id"`
    Room          Room      `json:"room"`
    CheckIn       time.Time `json:"check_in" validate:"required"`
    CheckOut      time.Time `json:"check_out" validate:"required"`
    TotalGuests   int       `json:"total_guests" validate:"required"`
    PaymentStatus string    `json:"payment_status"`
    Payments      []Payment `json:"payments"`
}


// Promotion model
type Promotion struct {
    gorm.Model
    Code        string    `json:"code" validate:"required"`
    Description string    `json:"description"`
    ValidFrom   time.Time `json:"valid_from" validate:"required"`
    ValidTo     time.Time `json:"valid_to" validate:"required"`
}

// Payment model
type Payment struct {
    gorm.Model
    BookingID     uint    `json:"booking_id" validate:"required"`
    Amount        float64 `json:"amount" validate:"required,gt=0"`
    PaymentMethod string  `json:"payment_method" validate:"required"`
    TransactionID string  `json:"transaction_id"`
}

// Employee model
type Employee struct {
    gorm.Model
    Name     string `json:"name" validate:"required"`
    Position string `json:"position"`
}

// Location model
type Location struct {
    gorm.Model
    Name      string  `json:"name" validate:"required"`
    Address   string  `json:"address"`
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}

// Content model
type Content struct {
    gorm.Model
    Title string `json:"title" validate:"required"`
    Body  string `json:"body" validate:"required"`
}
