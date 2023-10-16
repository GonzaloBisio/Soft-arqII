package main

import (
	hotelc "hotel-api/controllers"
)

func mapUrls() {

	routerUsuario := router.Group("/usuario")
	routerUsuario.Use(TokenMiddleware())

	routerAdmin := router.Group("/admin")
	routerAdmin.Use(AdminTokenMiddleware())

	//Hotel
	router.GET("/hotelId/:id", hotelc.GetHotelbyid)
	routerAdmin.POST("/insertHotel/:name/:Nroom/:descr", hotelc.InsertHotel)

	//User
	routerAdmin.GET("/userId/:id", userc.GetUserById)
	router.POST("/addUsuario/:name/:LastName/:DNI/:Password/:Email/:Admin", userc.AddUser)
	routerAdmin.GET("/users", userc.GetUsers)
	router.POST("/login", userc.Login)

	//Reservation
	routerAdmin.POST("/agregarReservation/:idHotel/:inicio/:final/:idUser/:habitacion", resrc.NewReserva)
	routerUsuario.GET("/reserva/:id", resrc.GetReservaById)
	routerAdmin.GET("/reservas", resrc.GetReservas)
	router.GET("/dispoibilidadDeReserva/:idHotel/:inicio/:final/:idUser/:habitacion", resrc.Dispoibilidad_de_reserva)
	routerUsuario.GET("/reservaByUserId/:user_id", resrc.GetReservasByUserId)
	//Amenmitie
	router.POST("/insertAmenitie/:name", amenc.InsertAmenitie)
	router.GET("/getAmenitie/:id", amenc.GetAmenitieById)
	router.GET("/getAmenities", amenc.GetAmenities)

	log.Info("Urls Cargadas")
}
