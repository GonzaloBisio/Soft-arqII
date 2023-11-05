package dtos

type userLogin struct{
	Email    string `json:"email"`
	Password string `json:"password"`
}