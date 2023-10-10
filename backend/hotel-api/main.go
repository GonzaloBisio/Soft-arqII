package main

import (
	"hotel-api/dao"
	"log"

	"github.com/streadway/amqp"
)

var rabbitMQConn *amqp.Connection
var rabbitMQCh *amqp.Channel

func init() {
	dao.InitializeMongoClient()
	rabbitMQURL := "amqp://tu_usuario:tu_contraseña@localhost:5672/"
	var err error
	rabbitMQConn, err = amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatal(err)
	}

	// Abre un canal de RabbitMQ
	rabbitMQCh, err = rabbitMQConn.Channel()
	if err != nil {
		log.Fatal(err)
	}

}
