package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type HotelService struct {
    Collection *mongo.Collection
}

func (s *HotelService) CreateHotel(ctx context.Context, hotel *models.Hotel) error {
    // Lógica para crear un hotel en MongoDB y notificar a RabbitMQ
}

func (s *HotelService) GetHotelByID(ctx context.Context, id string) (*models.Hotel, error) {
    // Lógica para obtener un hotel por ID
}

func (s *HotelService) UpdateHotel(ctx context.Context, id string, hotel *models.Hotel) error {
    // Lógica para modificar un hotel en MongoDB y notificar a RabbitMQ
}

// ... otros métodos de servicio
