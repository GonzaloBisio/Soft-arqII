package dtos

type ReservationCreateDto struct {
	UserId      int    `json:"user_id"`
	HotelId     string `json:"hotel_id"`
	InitialDate string `json:"initial_date"`
	FinalDate   string `json:"final_date"`
}
