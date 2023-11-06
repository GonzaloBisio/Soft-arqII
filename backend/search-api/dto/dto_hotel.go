package dto

type HotelDto struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type HotelDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AvailabilityResponse struct {
	Status bool `json:"ok_to_book"`
}

type HotelsDto []HotelDto
