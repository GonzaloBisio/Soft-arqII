package model

import "time"

type Reservation struct {
	Id          int       `gorm:"primaryKey"`
	InitialDate time.Time `gorm:"type:date"`
	FinalDate   time.Time `gorm:"type:date"`
	UserID      int       `gorm:"type:varchar(10);not null"`
	HotelID     string    `gorm:"type:varchar(10);not null"`
	User        string    `gorm:"type:varchar(10);not null"`
}

type Reservations []Reservation
