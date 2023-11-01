package dtos

type userDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UserName string `json:"username"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Type     bool   `json:"type"`
}

type usersDto []userDto