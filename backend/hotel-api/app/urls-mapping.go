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
	router.POST("/insertHotel", controllers.CreateHotel, controllers.PublishMessagePost)
	router.GET("/getHotels", controllers.GetHotels)
	router.POST("DeleteHotel/:id", controllers.PublishMessageDelete) //CRear delete de hotel
	router.POST("/UpdateHotel", controllers.UpdateHotel, controllers.PublishMessageUpdate)
	log.Info("Urls Cargadas")
}
