package app

import (
	log "github.com/sirupsen/logrus"
	resrc "user-reserva-disponibilidad-api/controllers"

	userController "user-reserva-disponibilidad-api/controllers"
)

func mapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/getUsers", userController.GetUsers)
	router.POST("/insertUser", userController.UserInsert) // Sign In

	//Reservation
	router.GET("/check", resrc.CheckStatus)
	router.POST("/reserva", resrc.NewReserva)
	router.GET("/reserva/:id", resrc.GetReservaById)
	router.GET("/reservas", resrc.GetReservas)
	router.GET("/reservaByUserId/:user_id", resrc.GetReservasByUserId)

	log.Info("Urls Cargadas")
}
