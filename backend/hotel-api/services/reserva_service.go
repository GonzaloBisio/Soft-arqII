package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"hotel-api/models"
	"hotel-api/utils/errors"
)

type reservaService struct {
	Collection *mongo.Collection
}

func (s *hotelService) GetReservaById(id string) (models.Reserva, errors.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (s *hotelService) InsertReserva(hotel models.Reserva) (models.Reserva, errors.ApiError) {
	//TODO implement me
	panic("implement me")
}

type reservaServiceInterface interface {
	GetReservaById(id string) (models.Reserva, errors.ApiError)
	InsertReserva(hotel models.Reserva) (models.Reserva, errors.ApiError)
}

var (
	ReservaService reservaServiceInterface
)

/*func init() {
	ReservaService = &reservaService{}
}
*/
