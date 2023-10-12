package daos

import (
	"search-api/dtos"
	"search-api/models"
)


type hotelService struct{}

type hotelServiceInterface interface {
    GetHotel(id string) (dtos.HotelDto, e.ApiError)
    CreateHotel(hotel models.Hotel) (dtos.HotelDto, e.ApiError)
    UpdateHotel(hotel models.Hotel) (dtos.HotelDto, e.ApiError)
    GetAll() (dtos.HotelsDto, e.ApiError)
    GetByCity(city string) (dtos.HotelsDto, e.ApiError)
    GetByAvailability(city string, checkIn string, checkOut string) (dtos.HotelsDto, e.ApiError)
}


var(
	HotelService hotelServiceInterface
)

func init(){
	HotelService = &hotelService{}
}

