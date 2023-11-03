package connection

import (
	"encoding/json"
	"log"
	"os"
	"search-api/config"
	"search-api/controller"
	dto "search-api/dto"

	"github.com/streadway/amqp"
)

var QueueConn *amqp.Connection

func handleError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}

}

func QueueConnection() {
	log.Print("Queue")
	QueueConn, err := amqp.Dial(config.RabitURL)
	handleError(err, "Can't connect to AMQP")
	defer QueueConn.Close()

	amqpChannel, err := QueueConn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("hotel", true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")

	err = amqpChannel.Qos(1, 0, false) //un mensjae a la vez
	handleError(err, "Could not configure QoS")

	messageChannel, err := amqpChannel.Consume( //Escuchando mensajes en esta cola
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)

			var queueDto dto.QueueDto

			err := json.Unmarshal(d.Body, &queueDto)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			log.Printf("ID %s, Action %s", queueDto.Id, queueDto.Action)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}

			if queueDto.Action == "INSERT" || queueDto.Action == "UPDATE" {
				err := controller.AddFromId(queueDto.Id)

				if err != nil {
					handleError(err, "Error inserting or deleting from Solr")
				}

			} else if queueDto.Action == "DELETE" {
				err := controller.Delete(queueDto.Id)

				if err != nil {
					handleError(err, "Error inserting or deleting from Solr")
				}
			}

		}
	}()

	// Stop for program termination
	<-stopChan

}
