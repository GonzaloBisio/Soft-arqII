package dto

type QueueDto struct {
	Id     string `json:"id"`
	Action string `json:"action"`
}

type DocDto struct {
	Doc HotelDTO `json:"doc"`
}
type AddDto struct {
	Add DocDto `json:"add"`
}
