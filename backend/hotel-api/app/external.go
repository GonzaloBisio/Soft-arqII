package app

import (
	"log"

	"github.com/streadway/amqp"
)

var rabbitMQCh *amqp.Channel
var rabbitMQConn *amqp.Connection

func init() {

	// Inicializa el cliente MongoDB
	InitializeMongoClient()

	// Configura la conexión a RabbitMQ
	rabbitMQURL := "amqp://user:password@localhost:5672/"
	var err error
	rabbitMQConn, err = amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatal(err)

	} else {
		log.Println("Rabit cargado")
	}

	// Abre un canal de RabbitMQ
	rabbitMQCh, err = rabbitMQConn.Channel()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Rabit Canal creado")
	}

	// Declara las colas de RabbitMQ para la creacion de hoteles
	queueName := "hotel_creation"
	_, err = rabbitMQCh.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

}
