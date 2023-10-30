package app

import (
	controllers "hotel-api/controllers"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	//Reservas
	router.GET("/ReservaId/:id", controllers.GetReservaById)
	router.POST("/insertReserva", controllers.CreateReserva)

	//Hotel
	router.GET("/hotelId/:id", controllers.GetHotelByID)
	router.POST("/insertHotel", controllers.CreateHotel)
	router.GET("/getHotels", controllers.GetHotels)

	log.Info("Urls Cargadas")
}
