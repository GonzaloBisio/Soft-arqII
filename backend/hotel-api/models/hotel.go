package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Reserva struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FechaIni time.Time          `bson:"fechaIni" json:"fechaIni"`
	FechaFin time.Time          `bson:"fechaFin" json:"fechaFin"`
}
type Hotel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Fotos       []string           `bson:"fotos" json:"fotos"`
	Amenities   []string           `bson:"amenities" json:"amenities"`
}

type Hotels struct {
	Hotels []Hotel `json:"hotels"`
}
