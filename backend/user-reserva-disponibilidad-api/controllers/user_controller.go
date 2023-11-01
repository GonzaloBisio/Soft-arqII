package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)


func GetUserById (c *gin.Context){
	log.Debug("User id to load: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))

	var UserDto users_dto.userDto
	UserDto, err := services.userService.GetUserById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, UserDto)
}

func GetUsers (c *gin.Context){
	var UsersDto users_dto.userDto 
	UsersDto, err := services.userService.GetAllUsers()

	if err != nil {
		c.JSON(err.Status(), err)
		return 
	}

	c.JSON(http.StatusOK, UsersDto)
}

func InsertUser (c *gin.Context){
	var UserDto users_dto.UserDto 
	err := c.BindJSON(&UserDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return 
	}

	UserDto, er := services.userService.InsertUser(UserDto)
	if err != nil {
		c.JSON(er.Status(), er)
		return 
	}
	c.JSON(http.StatusCreated, UserDto)
}

