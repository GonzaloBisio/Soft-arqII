package dtos

type hotelDto struct {
	Id               int    `json:"id"`
	HotelName        string `json:"hotel_name"`
	IdMongo          string `json:"id_mongo"`
	IdAmadeus        string `json:"id_amadeus"`
}

type hotelsDto []hotelDto


