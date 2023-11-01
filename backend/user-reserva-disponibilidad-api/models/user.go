package models

type User struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(250);not null"`
	LastName string `gorm:"type:varchar(250);not null"`
	Email    string `gorm:"type:varchar(50);not null;unique"`
	Password string `gorm:"type:varchar(150);not null"`
	Admin    int

	Bookings Bookings `gorm:"foreignKey:UserId"`
}

type Users []User
