package main

import (
	_ "github.com/go-sql-driver/mysql"
	"user-reserva-disponibilidad-api/app"
)

func main() {
	app.StartApp()
}
