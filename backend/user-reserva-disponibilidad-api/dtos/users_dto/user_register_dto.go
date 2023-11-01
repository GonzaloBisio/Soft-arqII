package dtos

type userDtoRegister struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type usersDtoRegister []userDtoRegister