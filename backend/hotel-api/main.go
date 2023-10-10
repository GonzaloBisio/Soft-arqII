package main

import (
	"hotel-api/dao"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

var rabbitMQConn *amqp.Connection
var rabbitMQCh *amqp.Channel

func init() {
    // Inicializa el cliente MongoDB
    dao.InitializeMongoClient()

    // Configura la conexión a RabbitMQ
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

    // Declara las colas de RabbitMQ según tus necesidades
    queueName := "hotel_creation"
    _, err = rabbitMQCh.QueueDeclare(queueName, false, false, false, false, nil)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    r := mux.NewRouter()

    if err := http.ListenAndServe(":3000", r); err != nil {
        log.Fatal(err)
    }
}
