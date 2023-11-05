package dtos

type checkDisponibilidad struct{
	StartDate int `json:"Start Date"`
	EndDate int `json:"End Date"`
}

type checkDisponibilidades []checkDisponibilidad