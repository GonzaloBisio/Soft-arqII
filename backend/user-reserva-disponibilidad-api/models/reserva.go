package models

import "time"

type Booking struct {
	Id          int       `gorm:"primaryKey"`
	InitialDate time.Time `gorm:"column:initial_date"`
	FinalDate   time.Time `gorm:"column:final_date"`
	User        string    `gorm:"type:varchar(10)"`
	UserId      string    `gorm:"column:user_id"`
	HotelId     string    `gorm:"column:hotel_id"`
}

type Bookings []Booking
