package queue

import (
	"log"
	"net/http"

	"hotel-api/utils/errors"

	"github.com/streadway/amqp"
)

type RabbitMQConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

// Representa una cola de RabbitMQ
type RabbitMQQueue struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	config  RabbitMQConfig
}

func NewRabbitMQQueue(config RabbitMQConfig) (*RabbitMQQueue, errors.ApiError) {
	conn, err := amqp.Dial("amqp://" + config.Username + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/")
	if err != nil {
		return nil, errors.NewRabbitMQError("Error al conectar a RabbitMQ", http.StatusInternalServerError, errors.CauseList{err.Error()})
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, errors.NewRabbitMQError("Error al abrir el canal de RabbitMQ", http.StatusInternalServerError, errors.CauseList{err.Error()})
	}

	return &RabbitMQQueue{
		conn:    conn,
		channel: channel,
		config:  config,
	}, nil
}

func (q *RabbitMQQueue) PublishMessage(queueName, message string) errors.ApiError {
	_, err := q.channel.QueueDeclare(
		queueName, // Nombre de la cola
		false,     // Durable
		false,     // AutoDelete
		false,     // Exclusive
		false,     // NoWait
		nil,       // Args
	)
	if err != nil {
		return errors.NewRabbitMQError("Error al declarar la cola de RabbitMQ", http.StatusInternalServerError, errors.CauseList{err.Error()})
	}

	// Publica el mensaje en la cola
	err = q.channel.Publish(
		"",        // Exchange
		queueName, // Queue
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return errors.NewRabbitMQError("Error al publicar el mensaje en la cola de RabbitMQ", http.StatusInternalServerError, errors.CauseList{err.Error()})
	}

	log.Printf("Mensaje enviado a la cola %s: %s", queueName, message)
	return nil
}

// Close cierra la conexión y el canal de RabbitMQ.
func (q *RabbitMQQueue) Close() {
	if q.channel != nil {
		q.channel.Close()
	}
	if q.conn != nil {
		q.conn.Close()
	}
}
