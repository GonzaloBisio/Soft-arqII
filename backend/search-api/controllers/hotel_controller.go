package controllers

import (
	"net/http"
	"search-api/dtos"
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
	var hotel dtos.HotelDto
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
	updatedHotel, err := services.HotelService.GetHotelByID(hotelID)
	if err != nil {
		c.JSON(err.Status(), err)
		return 
	}

	var hotelDto dtos.HotelDto

	if err := c.ShouldBindJSON(&hotelDto); err != nil {
		apiErr := e.NewBadRequestApiError("Pedido no valido")
		c.JSON(apiErr.Status(), apiErr)
		return 
	}

	updatedHotel.Name = hotelDto.Name
	updatedHotel.Description = hotelDto.Description

	_, err = services.HotelService.UpdateHotel(updatedHotel)
	if err != nil {
		apiErr := e.NewBadRequestApiError("Error al actualizar el Hotel")
		c.JSON(apiErr.Status(), apiErr)
		return 
	}

	c.JSON(http.StatusOK, updatedHotel)
}


/*
func GetHotelsByCity(c *gin.Context) {
	city := c.Param("city")
	hotelsDto, err := services.HotelService.GetHotelsByCity(city)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, hotelsDto)
	return
}
*/