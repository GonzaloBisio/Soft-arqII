package queue

import (
	"log"

	"github.com/streadway/amqp"
)

// la configuracion de rabbit
type RabbitMQConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

// esto representa lo que seria una ocla de RabbitMQ
type RabbitMQQueue struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	config  RabbitMQConfig
}

// instancia de RabbitMQQueue.
func NewRabbitMQQueue(config RabbitMQConfig) (*RabbitMQQueue, error) {
	conn, err := amqp.Dial("amqp://" + config.Username + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/")
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQQueue{
		conn:    conn,
		channel: channel,
		config:  config,
	}, nil
}

func (q *RabbitMQQueue) PublishMessage(queueName, message string) error {
	_, err := q.channel.QueueDeclare(
		queueName, // Nombre de la cola
		false,     // Durable
		false,     // AutoDelete
		false,     // Exclusive
		false,     // NoWait
		nil,       // Args
	)
	if err != nil {
		return err
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
		return err
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
