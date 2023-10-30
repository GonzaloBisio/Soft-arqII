package app

import (
	hotController "search-api/controllers"
)

func mapUrls() {
	router.GET("/hotel", hotController.GetAllHotels)
	router.GET("/hotel/:id", hotController.GetHotelByID)
	router.POST("/hotel", hotController.CreateHotel)
}
