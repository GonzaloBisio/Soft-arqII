package controllers

import (
	"hotel-api/dto"
	"hotel-api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Configura la conexión a RabbitMQ

func CreateHotel(c *gin.Context) {
	log.Println("Hotel registrado exitosamente")
	var newHotel dto.HotelDTO

	if err := c.ShouldBindJSON(&newHotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar la solicitud"})
		return
	}

	newHotel, err := services.HotelService.InsertHotel(newHotel)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.AddParam("id", newHotel.ID)
	c.JSON(http.StatusOK, newHotel)
}

func GetHotelByID(c *gin.Context) {
	hotelID := c.Param("id")
	var hotel dto.HotelDTO

	hotel, err := services.HotelService.GetHotelById(hotelID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotel)
}

func UpdateHotel(c *gin.Context) {
	var updatedHotel dto.HotelDTO
	if err := c.ShouldBindJSON(&updatedHotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar la solicitud"})
		return
	}

	updatedHotel, err := services.HotelService.UpdateHotel(updatedHotel)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, updatedHotel)
}

func GetHotels(c *gin.Context) {
	hotels, err := services.HotelService.GetHotels()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotels)
}

func DeleteHotelById(c *gin.Context) {
	hotelID := c.Param("id")

	err := services.HotelService.DeleteHotelById(hotelID)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	obj := map[string]string{

		"id":     hotelID,
		"Status": "DELETED",
	}

	c.AddParam("id", hotelID)
	c.JSON(http.StatusOK, obj)
}
