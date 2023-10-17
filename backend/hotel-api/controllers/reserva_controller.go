package controllers

import (
	"github.com/gin-gonic/gin"
	"hotel-api/models"
	"hotel-api/services"
	"hotel-api/utils/queue"
	"log"
	"net/http"
)

func CreateReserva(c *gin.Context) {
	log.Println("Reserva registrada exitosamente")
	var newReserva models.Reserva
	if err := c.ShouldBindJSON(&newReserva); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar la solicitud"})
		return
	}

	/*createdReserva, err := services.HotelService.InsertReserva(newReserva)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}*/

	rabbitMQ, err := queue.NewRabbitMQQueue(RabbitMQConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al configurar RabbitMQ"})
		return
	}
	defer rabbitMQ.Close()

	/*message := "Se creó una nueva reserva: " + createdReserva.FechaIni.Format("2006-01-02") + " - " + createdReserva.FechaFin.Format("2006-01-02")
		queueName := "reserva_creation"
		err = rabbitMQ.PublishMessage(queueName, message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar mensaje a la cola RabbitMQ"})
			return
		}

		c.JSON(http.StatusCreated, createdReserva)
	}*/
}
func GetReservaById(c *gin.Context) {
	ReservaId := c.Param("id")

	reserva, err := services.HotelService.GetHotelById(ReservaId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reserva)
}
