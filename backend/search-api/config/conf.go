package config

import "fmt"

var (
	///////Zona de Ssolr/////
	solRhost       = "localhost"
	solRPort       = 8983
	SolRCollection = "hotels"

	SolrURL = fmt.Sprintf("http://%s:%d", solRhost, solRPort)

	////Zona de RabbitMQ/////
	rabbithost     = "localhost"
	rabbitPort     = 5672
	rabbitPaswword = "password"
	rabbitUser     = "user"

	RabitURL = fmt.Sprintf("amqp://%s:%s@%s:%d/", rabbitUser, rabbitPaswword, rabbithost, rabbitPort)

	////ZONA DE HOTEL/////
	hotelHost = "localhost"
	hotelPort = 8000

	HotelUrl = fmt.Sprintf("http://%s:%d/", hotelHost, hotelPort)

	/////Zona de Reservations////
	reservationHost = "localhost"
	reservationPort = 8003

	ReservationlUrl = fmt.Sprintf("http://%s:%d", hotelHost, hotelPort)
)
