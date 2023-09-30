package space_trips_system

import "time"

type User struct {
	Id       int     `json:"-"`
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Phone    string  `json:"phone"`
	Balance  float64 `json:"balance"`
}

type UserTrips struct {
	Id           int       `json:"id"`
	PurchaseDate time.Time `json:"purchaseDate"`
	UserId       int       `json:"userId"`
	TripId       int       `json:"tripId"`
}
