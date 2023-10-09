package services

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelService struct {
	Collection *mongo.Collection
}

type hotelServiceInterface interface {
	GetHotels()(dtos.Hotels, e.ApiError)
	GetHotelById(id string) (dtos.Hotel, e.ApiError)
	InsertHotel(Hotel dtos.Hotel) (dtos.Hotel, e.ApiError)
	UpdateHotel(Hotel dtos.Hotel) (dtos.Hotel, e.ApiError)
}

var(
	hotelService hotelServiceInterface
)

func init() {
	hotelService = &HotelService{}
}

func (s *HotelService) GetHotels()(dtos.Hotels, e.ApiError) { //Para arreglar esto, tenemos que, importar los dtos que
	//no me esta saliendo, y ademas, implmentar utils del proyecto pasado con los errores, por eso e.ApiError, etc. 
	//Dsps, creo que la logica esta bien.
	//Ademas, el hotelDao, significa que 
	hotels, err := hotelDao.Client.GetAll()
	if err != nil {
		return dtos.Hotels{}, e.NewInternalServerApiError("Ningun hotel encontrado", err)
	}

	//si no hay error, pasamos los datos al dto
	var hotelDtos = make([]dtos.Hotel, 0)
	for _, hotel := range hotels {
		hotelDto := dtos.Hotel{
			ID: hotel.ID.hex(),
			Name: hotel.Name,
			Description: hotel.Description,
		}
		hotelDtos = append(hotelDtos, hotelDto)
	}

	final := dtos.Hotels{
		Hotels: hotelDtos,
	}

	return final, nil
}

func (s *HotelService) GetHotelById (id string)(dtos.Hotel, e.ApiError){
	hotel, err := hotelDao.Client.GetHotelById(id)
	if err != nil {
		return dtos.Hotel{}, e.NewInternalServerApiError("Ningun hotel existente con ese id", err)
	}

	hotelDto := dtos.Hotel{
		ID: hotel.ID.hex(),
		Name: hotel.Name,
		Description: hotel.Description,
	}

	return hotelDto, nil
}

