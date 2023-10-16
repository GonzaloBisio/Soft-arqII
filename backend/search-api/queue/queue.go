package queue

import (
	"encoding/json"
	"log"
	hotController "search-api/controllers"
	"search-api/dtos"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func StartReceiving() {
	var err error
	conn, err = amqp.Dial("amqp://user:password@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err = ch.QueueDeclare(
		"hotels", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// Decodificar el mensaje en un DTO de hotel
			var hotelDto dtos.HotelDto
			err := json.Unmarshal(d.Body, &hotelDto)
			if err != nil {
				log.Printf("Error al decodificar el mensaje: %v", err)
				continue
			}

			// Intentar crear o actualizar el hotel
			if hotelDto.ID == "" {
				// Si el hotel no tiene un ID, crea un nuevo hotel
				createdHotel, err := hotController.CreateHotel(hotelDto)
				if err != nil {
					log.Printf("Error al crear el hotel: %v", err)
				} else {
					log.Printf("Hotel creado con éxito: %v", createdHotel)
				}
			} else {
				// Si el hotel tiene un ID, actualiza un hotel existente
				updatedHotel, err := hotController.GetHotelByID(hotelDto.ID)
				if err != nil {
					log.Printf("Error al obtener el hotel existente: %v", err)
				} else {
					// Actualizar el hotel con los datos del DTO
					updatedHotel.Name = hotelDto.Name
					updatedHotel.City = hotelDto.City
					updatedHotel.Description = hotelDto.Description
					updatedHotel.Thumbnail = hotelDto.Thumbnail
					updatedHotel.Images = hotelDto.Images
					updatedHotel.Amenities = hotelDto.Amenities

					_, err = hotController.UpdateHotel(updatedHotel)
					if err != nil {
						log.Printf("Error al actualizar el hotel: %v", err)
					} else {
						log.Printf("Hotel actualizado con éxito: %v", updatedHotel)
					}
				}
			}
		}
	}()
	log.Printf("Subscripción a la cola con éxito")
	<-forever
}
