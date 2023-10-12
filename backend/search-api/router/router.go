package router

import (
	hotController "search-api/controllers"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	router.GET("/hotel", hotController.GetAllHotels)
	router.GET("/hotel/:id", hotController.GetHotelByID)
	router.POST("/hotel", hotController.CreateHotel)
}