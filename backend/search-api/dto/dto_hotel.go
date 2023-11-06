package dto

type HotelDto struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type HotelDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Country     string `json:"country"`
	City        string `json:"city"`
}

type HotelsDto struct {
	Hotels []HotelDTO `json:"hotels"`
}
