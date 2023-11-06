package dto

type HotelDTO struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Fotos       []string `json:"fotos"`
	Amenities   []string `json:"amenities"`
	Country     string   `json:"country"`
	City        string   `json:"city"`
}

type HotelsDto struct {
	Hotels []HotelDTO `json:"hotels"`
}
