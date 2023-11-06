package dtos

type ReservaDto struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	HotelId   string `json:"hotel_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type reservasDto []ReservaDto
