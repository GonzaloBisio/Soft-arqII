package config

import "fmt"

var (
	solRhost       = "localhost"
	solRPort       = 8983
	SolRCollection = "Hotels"

	SolrURL = fmt.Sprintf("http://%s:%d/solr/", solRhost, solRPort)

	rabbithost     = "localhost"
	rabbitPort     = 5672
	rabbitPaswword = "password"
	rabbitUser     = "user"

	RabitURL = fmt.Sprintf("amqp://%s:%s@%s:%d/", rabbitUser, rabbitPaswword, rabbithost, rabbitPort)

	hotelHost = "localhost"
	hotelPort = 8000

	HotelUrl = fmt.Sprintf("http://%s:%d/", hotelHost, hotelPort)
)
