package dtos

type HotelDto struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    City        string `json:"city"`
}

type HotelsDto struct {
    Hotels []HotelDto `json:"hotels"`
}