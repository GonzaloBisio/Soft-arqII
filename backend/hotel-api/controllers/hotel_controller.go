package controllers

import (
	"hotel-api/models"
	"hotel-api/services"
	"hotel-api/utils/queue"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHotel(c *gin.Context) {
    var newHotel models.Hotel
    if err := c.ShouldBindJSON(&newHotel); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar la solicitud"})
        return
    }

    createdHotel, err := services.HotelService.InsertHotel(newHotel) 
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }

    // Configura la conexión a RabbitMQ
    rabbitMQConfig := queue.RabbitMQConfig{
        Username: "tu_usuario",
        Password: "tu_contraseña",
        Host:     "localhost",
        Port:     "5672",
    }
    rabbitMQ, err := queue.NewRabbitMQQueue(rabbitMQConfig)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al configurar RabbitMQ"})
        return
    }
    defer rabbitMQ.Close()

    //Este es el mensaje que se envia a la cola
    message := "Se creó un nuevo hotel: " + createdHotel.Name
    // Nombre de la cola
    queueName := "hotel_creation"
    // Envia el mensaje a la cola
    err = rabbitMQ.PublishMessage(queueName, message)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar mensaje a la cola RabbitMQ"})
        return
    }

    c.JSON(http.StatusCreated, createdHotel)
}

func GetHotelByID(c *gin.Context) {
    hotelID := c.Param("id")

    hotel, err := services.HotelService.GetHotelByID(hotelID)
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }

    c.JSON(http.StatusOK, hotel)
}

func UpdateHotel(c *gin.Context) {
    var updatedHotel models.Hotel
    if err := c.ShouldBindJSON(&updatedHotel); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar la solicitud"})
        return
    }

    updatedHotel, err := services.HotelService.UpdateHotel(updatedHotel)
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }

    rabbitMQConfig := queue.RabbitMQConfig{
        Username: "tu_usuario",
        Password: "tu_contraseña",
        Host:     "localhost",
        Port:     "5672",
    }
    rabbitMQ, err := queue.NewRabbitMQQueue(rabbitMQConfig)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al configurar RabbitMQ"})
        return
    }
    defer rabbitMQ.Close()

    message := "Se actualizó un hotel: " + updatedHotel.Name
    queueName := "hotel_update"
    
    err = rabbitMQ.PublishMessage(queueName, message)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar mensaje a la cola RabbitMQ"})
        return
    }

    c.JSON(http.StatusOK, updatedHotel)
}
