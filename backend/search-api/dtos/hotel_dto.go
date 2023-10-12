package dtos

type HotelDto struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	City         string   `json:"city"`
	Description  string   `json:"description"`
	Thumbnail    string   `json:"thumbnail"`
	Images       []string `json:"images"`
	Amenities    []string `json:"amenities"`
	Availability bool     `json:"availability"`
}

type HotelsDto struct {
	Hotels []HotelDto `json:"hotels"`
}