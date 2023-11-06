package app

import (
	log "github.com/sirupsen/logrus"
	resrc "user-reserva-disponibilidad-api/controllers"
)

func mapUrls() {

	//User
	/*	router.GET("/userId/:id", userc.GetUserById)
		router.POST("/addUsuario/:name/:LastName/:DNI/:Password/:Email/:Admin", userc.AddUser)
		routerAdmin.GET("/users", userc.GetUsers)
		router.POST("/login", userc.Login)*/

	//Reservation
	router.POST("/reserva", resrc.NewReserva)
	router.GET("/reserva/:id", resrc.GetReservaById)
	router.GET("/reservas", resrc.GetReservas)
	router.GET("/reservaByUserId/:user_id", resrc.GetReservasByUserId)

	log.Info("Urls Cargadas")
}
