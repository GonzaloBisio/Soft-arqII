package model

import "time"

type Reservation struct {
	Id          int       `gorm:"primaryKey"`
	InitialDate time.Time `gorm:"column:initial_date;not null"`
	FinalDate   time.Time `gorm:"column:final_date;not null"`
	UserID      int
	HotelID     string
	User        User `gorm:"foreignKey:UserID"`
}

type Reservations []Reservation
