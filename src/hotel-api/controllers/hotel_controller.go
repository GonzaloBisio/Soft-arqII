package controllers

import (
	"hotel-api/models"
	"hotel-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateHotel crea un nuevo hotel.
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

    c.JSON(http.StatusCreated, createdHotel)
}

// GetHotelByID obtiene un hotel por ID.
func GetHotelByID(c *gin.Context) {
    hotelID := c.Param("id")

    hotel, err := services.HotelService.GetHotelByID(hotelID)
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }

    c.JSON(http.StatusOK, hotel)
}

// UpdateHotel actualiza un hotel.
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

    c.JSON(http.StatusOK, updatedHotel)
}
