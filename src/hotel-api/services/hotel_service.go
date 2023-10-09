package services

import (
	"hotel-api/dao"
	"hotel-api/errors"
	"hotel-api/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type HotelService struct {
	Collection *mongo.Collection
}

type hotelServiceInterface interface {
	GetHotels()(models.Hotels, e.ApiError)
	GetHotelById(id string) (models.Hotel, e.ApiError)
	InsertHotel(Hotel models.Hotel) (models.Hotel, e.ApiError)
	UpdateHotel(Hotel models.Hotel) (models.Hotel, e.ApiError)
}

var(
	hotelService hotelServiceInterface
)

func init() {
	hotelService = &HotelService{}
}

func (s *HotelService) GetHotels() (models.Hotels, errors.ApiError) {
    hotels, err := dao.Client.GetAll()
    if err != nil {
        return models.Hotels{}, errors.NewInternalServerApiError("Ningun hotel encontrado", err)
    }

    var hotelDtos = make([]models.Hotel, 0)
    for _, hotel := range hotels {
        hotelDto := models.Hotel{
            ID:          hotel.ID.Hex(),
            Name:        hotel.Name,
            Description: hotel.Description,
        }
        hotelDtos = append(hotelDtos, hotelDto)
    }

    final := models.Hotels{
        Hotels: hotelDtos,
    }

    return final, nil
}

func (s *HotelService) GetHotelById(id string) (models.Hotel, errors.ApiError) {
    hotel, err := dao.Client.GetHotelById(id)
    if err != nil {
        return models.Hotel{}, errors.NewInternalServerApiError("Ningun hotel existente con ese id", err)
    }

    hotelDto := models.Hotel{
        ID:          hotel.ID.Hex(),
        Name:        hotel.Name,
        Description: hotel.Description,
    }

    return hotelDto, nil
}

func (s *HotelService) InsertHotel(hotel models.Hotel) (models.Hotel, errors.ApiError) {
    insertedHotel, err := dao.Client.Insert(hotel)
    if err != nil {
        return models.Hotel{}, errors.NewInternalServerApiError("Error al insertar el hotel en la base de datos", err)
    }

    hotelDto := models.Hotel{
        ID:          insertedHotel.ID.Hex(),
        Name:        insertedHotel.Name,
        Description: insertedHotel.Description,
    }

    return hotelDto, nil
}

func (s *HotelService) UpdateHotel(hotel models.Hotel) (models.Hotel, errors.ApiError) {
    updatedHotel, err := dao.Client.Update(hotel)
    if err != nil {
        return models.Hotel{}, errors.NewInternalServerApiError("Error al actualizar el hotel en la base de datos", err)
    }

    hotelDto := models.Hotel{
        ID:          updatedHotel.ID.Hex(),
        Name:        updatedHotel.Name,
        Description: updatedHotel.Description,
    }

    return hotelDto, nil
}

