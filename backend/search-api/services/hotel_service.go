package services

import (
	"search-api/daos"
	"search-api/dtos"
	e "search-api/errors"
	model "search-api/models"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetAllHotels() (dtos.HotelsDto, e.ApiError)
	GetHotelByID(id string) (dtos.HotelDto, e.ApiError)
	CreateHotel(hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError)
	UpdateHotel(hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError)
	//GetByCity(city string) (dtos.HotelsDto, e.ApiError)
	//GetByAvailability(city string, checkIn string, checkOut string) (dtos.HotelsDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetAllHotels() (dtos.HotelsDto, e.ApiError) {
	var hotelDtos dtos.HotelsDto
	hotelDtos.Hotels = []dtos.HotelDto{}
	hotelClient := daos.NewHotelSolrDAO()
	hotels, err := hotelClient.GetAllHotels()
	if err != nil {
		return hotelDtos, e.NewBadRequestApiError("Error al obtener hoteles")
	}

	for _, hotel := range hotels {
		hotelDto := dtos.HotelDto{
			ID:          hotel.ID,
			Name:        hotel.Name,
			City:        hotel.City,
			Description: hotel.Description,
			Thumbnail:   hotel.Thumbnail,
			Images:      hotel.Images,
			Amenities:   hotel.Amenities,
		}
		hotelDtos.Hotels = append(hotelDtos.Hotels, hotelDto)
	}

	return hotelDtos, nil
}

func (s *hotelService) GetHotelByID(id string) (dtos.HotelDto, e.ApiError) {
	hotelClient := daos.NewHotelSolrDAO()
	hotel, err := hotelClient.GetById(id)
	if err != nil {
		return dtos.HotelDto{}, e.NewNotFoundApiError("Ningún hotel encontrado con ese ID")
	}

	hotelDto := dtos.HotelDto{
		ID:          hotel.ID,
		Name:        hotel.Name,
		City:        hotel.City,
		Description: hotel.Description,
		Thumbnail:   hotel.Thumbnail,
		Images:      hotel.Images,
		Amenities:   hotel.Amenities,
	}

	return hotelDto, nil
}

func (s *hotelService) CreateHotel(hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError) {
	//hotelClient := daos.NewHotelSolrDAO()
	var hotel model.Hotel

	hotel.ID = hotelDto.ID
	hotel.City = hotelDto.City
	hotel.Description = hotelDto.Description
	hotel.Thumbnail = hotelDto.Thumbnail
	hotel.Amenities = hotelDto.Amenities
	hotel.Images = hotelDto.Images

	hotelClient := daos.NewHotelSolrDAO()
	err := hotelClient.CreateHotel(&hotel)

	if err != nil {
		return dtos.HotelDto{}, e.NewInternalServerApiError("Error al crear el hotel", err)
	}
	
	return hotelDto, nil
}

func (s *hotelService) UpdateHotel(hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError) {
    newHotel := model.Hotel{
        ID:          hotelDto.ID,
        Name:        hotelDto.Name,
        City:        hotelDto.City,
        Description: hotelDto.Description,
        Thumbnail:   hotelDto.Thumbnail,
        Images:      hotelDto.Images,
        Amenities:   hotelDto.Amenities,
    }

	hotelClient := daos.NewHotelSolrDAO()
    err := hotelClient.UpdateHotel(&newHotel)
    if err != nil {
        return dtos.HotelDto{}, e.NewInternalServerApiError("Error al actualizar el hotel", err)
    }

    return hotelDto, nil
}


//Aca faltaria hacer para obtener por ciudad
//y tambien habria que hacer unas funciones para checkear la disponibilidad segun fechas

