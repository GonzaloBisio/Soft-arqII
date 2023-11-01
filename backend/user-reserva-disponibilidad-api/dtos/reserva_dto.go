package dtos

type reservaDto struct{
	Id 			int `json: "id"`
	userId 		int `json:"user_id"`
	hotelId 	int `json:"hotel_id"`
	StartDate 	int `json: start_date`
	EndDate 	int `json:"end_date"`
}

type reservasDto []reservaDto