package services

import (
	"hotel-api/dao"
	"hotel-api/models"
	"hotel-api/utils/errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type hotelService struct {
	Collection *mongo.Collection
}

type hotelServiceInterface interface {
	GetHotels() (models.Hotels, errors.ApiError)
	GetHotelById(id string) (models.Hotel, errors.ApiError)
	InsertHotel(hotel models.Hotel) (models.Hotel, errors.ApiError)
	UpdateHotel(hotel models.Hotel) (models.Hotel, errors.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotels() (models.Hotels, errors.ApiError) {
	hotels, err := dao.Client.GetAll()
	if err != nil {
		return models.Hotels{}, errors.NewInternalServerApiError("Ningun hotel encontrado", err)
	}

	var hotelDtos = make([]models.Hotel, 0)
	for _, hotel := range hotels {
		hotelDto := models.Hotel{
			ID:          hotel.ID,
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

func (s *hotelService) GetHotelById(id string) (models.Hotel, errors.ApiError) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Hotel{}, errors.NewBadRequestApiError("no se que paso")
	}

	hotel, err := dao.Client.GetHotelById(objectID.Hex()) // Convierte ObjectID a cadena
	if err != nil {
		return models.Hotel{}, errors.NewBadRequestApiError("no se que paso")
	}

	return hotel, nil
}

func (s *hotelService) InsertHotel(hotel models.Hotel) (models.Hotel, errors.ApiError) {
	hotelInsertado, err := dao.Client.Insert(hotel)
	if err != nil {
		return hotel, errors.NewInternalServerApiError("Error al insertar el hotel en la base de datos", err)
	}
	return hotelInsertado, nil
}

func (s *hotelService) UpdateHotel(hotel models.Hotel) (models.Hotel, errors.ApiError) {

	_, err := dao.Client.Update(hotel)
	if err != nil {
		return models.Hotel{}, errors.NewInternalServerApiError("Error al actualizar el hotel en la base de datos", err)
	}
	return hotel, nil
}
