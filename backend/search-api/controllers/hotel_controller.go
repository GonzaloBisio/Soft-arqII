package controllers

import (
	"net/http"
	"search-api/models"
	"search-api/services"
	e "search-api/utils/errors"

	"github.com/gin-gonic/gin"
)

func GetAllHotels(c *gin.Context) {
	hotels, err := services.HotelService.GetAllHotels()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotels)
}

func GetHotelByID(c *gin.Context){
	hotelID := c.Param("id")
	hotel, err := services.HotelService.GetHotelByID(hotelID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotel)
}

func CreateHotel(c *gin.Context){
	var hotel models.Hotel
	if err := c.ShouldBindJSON(&hotel); err != nil {
		apiErr := e.NewBadRequestApiError("Datos invalidos")
		c.JSON(apiErr.Status(), apiErr)
		return 
	}

	createdHotel, err := services.HotelService.CreateHotel(hotel)
	if err != nil {
		c.JSON(err.Status(), err)
		return 
	}

	c.JSON(http.StatusOK, createdHotel)
}

func UpdateHotel(c *gin.Context) {
	hotelID := c.Param("id")
	var hotel models.Hotel
	if err := c.ShouldBindJSON(&hotel); err != nil {
		apiErr := e.NewBadRequestApiError("Datos inválidos")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	hotel.ID = hotelID
	updatedHotel, err := services.HotelService.UpdateHotel(hotel)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, updatedHotel)
}