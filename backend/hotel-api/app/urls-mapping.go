package app

import (
	controllers "hotel-api/controllers"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	/*
		routerUsuario := router.Group("/usuario")
		routerUsuario.Use(TokenMiddleware())

		routerAdmin := router.Group("/admin")
		routerAdmin.Use(AdminTokenMiddleware())
	*/

	//Reservas
	router.GET("/ReservaId/:id", controllers.GetReservaById)
	router.POST("/insertReserva", controllers.CreateReserva)

	//Hotel
	router.GET("/hotelId/:id", controllers.GetHotelByID)
	router.POST("/insertHotel", controllers.CreateHotel)

	log.Info("Urls Cargadas")
}
