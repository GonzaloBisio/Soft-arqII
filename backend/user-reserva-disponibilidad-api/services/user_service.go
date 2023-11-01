package services

/*
import (
	uClient "user-reserva-disponibilidad-api/dao"
	model "user-reserva-disponibilidad-api/models"
	e "user-reserva-disponibilidad-api/utils/errors/errors.go"
	//uDto "user-reserva-disponibilidad-api/dtos/users_dto"
	// uDto "user-reserva-disponibilidad-api/dtos/uDto/users_dto" siempre que guardo el archivo se me borra
)



type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (uDto.userDto, e.ApiError)
	GetUsers() (uDto.usersDto, e.ApiError)
	InsertUser(userDto uDto.userDtoRegister) (uDto.userDto, e.ApiError)
	UserLogin(userDto uDto.userLogin) (uDto.userLoginResponse, e.ApiError)
	IsEmailTaken(email string) bool
	DeleteUserById(id int) e.ApiError
	UpdateUser(userDto uDto.userDtoRegister, id int) e.ApiError
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

//Readaptada
func (s *userService) GetUserById(id int) (uDto.userDto, e.ApiError){
	var user model.User = uClient.GetUserById(id)
	var userDto uDto.userDto

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("User not found")
	}

	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.Email = user.Email
	userDto.Admin = user.Admin

	return userDto, nil
}

func (s *userService) GetUsers() (uDto.usersDto, e.ApiError){
	var users model.Users = uClient.GetAllUsers()
	var UsersDto uDto.usersDto

	for _, user := range users {
		var UserDto uDto.userDto

		//Si el user es Admin no lo incluye en el return

		if !UserDto.Type {
			UserDto.Name = user.Name
			UserDto.LastName = user.LastName
			UserDto.UserName = user.UserName
			UserDto.Address = user.Address
			UserDto.Email = user.Email
			UserDto.Id = user.Id
			UserDto.Type = user.Type
		}

		UsersDto = append(UsersDto, UserDto)
	}

	return UsersDto, nil
}

*/