package services

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelService struct {
	Collection *mongo.Collection
}

// ... otros métodos de servicio
