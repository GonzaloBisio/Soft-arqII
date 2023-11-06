package controller

import (
	"net/http"
	"search-api/dto"
	"search-api/services"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/log"
)

func AddFromId(id string) error {

	err := services.SolrService.AddFromId(id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFromId(id string) error {
	err := services.SolrService.DeleteFromId(id)
	if err != nil {
		return err
	}
	return nil
}

func GetQuery(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	query := c.Param("searchQuery")

	hotelsDto, err := services.SolrService.GetQuery(query, "fiueld")
	if err != nil {
		c.JSON(http.StatusBadRequest, hotelsDto)
		return
	}

	log.Debug(hotelsDto)

	c.JSON(http.StatusOK, hotelsDto)
}

//654513064fd3abb838f8c48d
