package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reserva struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FechaIni string             `bson:"fechaIni" json:"fechaIni"`
	FechaFin string             `bson:"fechaFin" json:"fechaFin"`
}
type Hotel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Fotos       []string           `bson:"fotos" json:"fotos"`
	Amenities   []string           `bson:"amenities" json:"amenities"`
	Reservas    []Reserva          `bson:"reservas" json:"reservas"`
}

type Hotels struct {
	Hotels []Hotel `json:"hotels"`
}
