package services

import (
	"search-api/daos"
	"search-api/dtos"
	e "serch-api/utils/errors" //no se porque me tira error en la importación.
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetAllHotels() (dtos.HotelsDto, e.ApiError)
	GetHotelByID(id string) (dtos.HotelDto, e.ApiError)
	CreateHotel(hotel dtos.HotelDto) (dtos.HotelDto, e.ApiError)
	UpdateHotel(hotel dtos.HotelDto) (dtos.HotelDto, e.ApiError)
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
	hotels, err := hotelClient.GetAll()
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
	hotel, err := hotelClient.Get(id)
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

func (s *hotelService) CreateHotel(hotel dtos.HotelDto) (dtos.HotelDto, e.ApiError) {
	hotelClient := daos.NewHotelSolrDAO()
	err := hotelClient.Create(&hotel)
	if err != nil {
		return dtos.HotelDto{}, e.NewInternalServerApiError("Error al crear el hotel", err)
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

func (s *hotelService) UpdateHotel(hotel dtos.HotelDto) (dtos.HotelDto, e.ApiError) {
	hotelClient := daos.NewHotelSolrDAO()
	err := hotelClient.Update(&hotel)
	if err != nil {
		return dtos.HotelDto{}, e.NewInternalServerApiError("Error al actualizar el hotel", err)
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

//Aca faltaria hacer para obtener por ciudad
//y tambien habria que hacer unas funciones para checkear la disponibilidad segun fechas

