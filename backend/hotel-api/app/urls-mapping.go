package app

import (
	hotelc "hotel-api/controllers"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	/*
		routerUsuario := router.Group("/usuario")
		routerUsuario.Use(TokenMiddleware())

		routerAdmin := router.Group("/admin")
		routerAdmin.Use(AdminTokenMiddleware())
	*/
	//Hotel
	router.GET("/hotelId/:id", hotelc.GetHotelByID)
	router.POST("/insertHotel", hotelc.CreateHotel)

	log.Info("Urls Cargadas")
}
