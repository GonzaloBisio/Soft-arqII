package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"search-api/dtos"
	e "search-api/errors"
	"search-api/services"

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
	updatedHotel.City = hotelDto.City
	updatedHotel.Description = hotelDto.Description
	updatedHotel.Thumbnail = hotelDto.Thumbnail
	updatedHotel.Images = hotelDto.Images
	updatedHotel.Amenities = hotelDto.Amenities

	_, err = services.HotelService.UpdateHotel(updatedHotel)
	if err != nil {
		apiErr := e.NewBadRequestApiError("Error al actualizar el Hotel")
		c.JSON(apiErr.Status(), apiErr)
		return 
	}

	c.JSON(http.StatusOK, updatedHotel)
}

func GetOrInsertByID(id string) {
	// Hacer una solicitud a hotel-api pidiendo todos los datos del hotel
	url := fmt.Sprintf("http://hotel-api:8000/hotel/%s", id)

	// Realizar la solicitud HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al hacer la solicitud HTTP:", err)
		return
	}
	defer resp.Body.Close()

	// Verificar si la respuesta fue exitosa (código 200)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("La solicitud no fue exitosa. Código de respuesta:", resp.StatusCode)
		return
	}

	// Leer el cuerpo de la respuesta HTTP
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta HTTP:", err)
		return
	}

	// Deserializar la respuesta en un objeto HotelDto
	var hotelResponse dtos.HotelDto
	if err := json.Unmarshal(body, &hotelResponse); err != nil {
		fmt.Println("Error al deserializar la respuesta:", err)
		return
	}

	// Comprobar si ya tienes cargado el hotel en Solr
	hotelSolr, err := services.HotelService.GetHotelByID(id)
	if err != nil {
		// Si no lo tienes cargado, entonces lo agregas
		_, err := services.HotelService.CreateHotel(hotelResponse)
		if err != nil {
			// Manejar el error de creación
			fmt.Println("Error al crear el hotel:", err)
			return
		}
		fmt.Println("Hotel nuevo agregado:", id)
		return
	}

	// Si ya lo tienes cargado, haces la actualización
	// Actualiza los campos del hotel existente con los nuevos valores
	hotelSolr.Name = hotelResponse.Name
	hotelSolr.Description = hotelResponse.Description
	hotelSolr.Thumbnail = hotelResponse.Thumbnail
	hotelSolr.Images = hotelResponse.Images
	hotelSolr.Amenities = hotelResponse.Amenities

	// Actualiza el hotel en Solr
	_, err = services.HotelService.UpdateHotel(hotelSolr)
	if err != nil {
		// Manejar el error de actualización
		fmt.Println("Error al actualizar el hotel:", err)
		return
	}
	return
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