package services

import (
	uClient "user-reserva-disponibilidad-api/dao"
	model "user-reserva-disponibilidad-api/models"
	e "user-reserva-disponibilidad-api/utils/errors/errors.go"
)



type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (users_dto.userDto, e.ApiError)
	GetUsers() (users_dto.usersDto, e.ApiError)
	InsertUser(userDto users_dto.userDtoRegister) (users_dto.userDto, e.ApiError)
	UserLogin(userDto users_dto.userLogin) (users_dto.userLoginResponse, e.ApiError)
	IsEmailTaken(email string) bool
	DeleteUserById(id int) e.ApiError
	UpdateUser(userDto users_dto.userDtoRegister, id int) e.ApiError
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

//Readaptada
func (s *userService) GetUserById(id int) (users_dto.userDto, e.ApiError){
	var user model.User = uClient.GetUserById(id)
	var userDto users_dto.userDto

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("User not found")
	}

	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.Email = user.Email
	userDto.Admin = user.Admin

	return userDto, nil
}

func (s *userService) GetUsers() (users_dto.usersDto, e.ApiError){
	var users model.User = uClient.GetAllUsers()
	usersDto := users_dto.usersDto{
		Users: make([]users_dto.usersDto, len(users)),
	}

	for i, user := range users {
		userDto:= users_dto.userDto{
			Id: user.Id, 
			Name: user.Name,
			LastName: user.LastName,
			Email: user.Email,
			Admin: user.Admin,
		}

		usersDto.Users[i] = userDto 
	}

	return usersDto, nil 
}

