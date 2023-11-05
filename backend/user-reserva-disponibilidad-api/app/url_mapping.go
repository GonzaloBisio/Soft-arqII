package app

import (
	log "github.com/sirupsen/logrus"
	hotelController "user-reserva-disponibilidad-api/controllers"
)

func mapUrls() {

	router.GET("/hotels", hotelController.GetHotels)
	router.POST("/hotels", hotelController.AddHotel)
	router.GET("/hotels/:id", hotelController.GetHotelByID)
	log.Info("Urls Cargadas")
}
