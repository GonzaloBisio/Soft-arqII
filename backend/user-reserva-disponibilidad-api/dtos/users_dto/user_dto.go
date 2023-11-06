package dtos

type UserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UserName string `json:"username"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Type     bool   `json:"type"`
	DNI      string `json:"dni"`
}

type UsersDto []UserDto
