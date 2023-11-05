package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	Hotel "user-reserva-disponibilidad-api/model"
)

var hoteles = []Hotel.Hotel{{ID: "1", Item: "Hotel 1", Completed: false},
	{ID: "2", Item: "Hotel 2", Completed: false},
	{ID: "3", Item: "Hotel 3", Completed: false},
}

func GetHotels(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, hoteles)
}

func GetHotelByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range hoteles {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Hotel no encontrado"})
}

func AddHotel(c *gin.Context) {
	var newHotel Hotel.Hotel

	if err := c.BindJSON(&newHotel); err != nil {
		return
	}

	hoteles = append(hoteles, newHotel)
	c.IndentedJSON(http.StatusCreated, newHotel)
}
