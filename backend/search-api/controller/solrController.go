package controller

import (
	"net/http"
	"search-api/dto"
	"search-api/services"

	"github.com/gin-gonic/gin"
	log "github.com/gofiber/fiber/v2/log"
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

func SerchQuery(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	query := c.Param("searchQuery")

	hotelsDto, err := services.SolrService.SerchQuery(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, hotelsDto)
		return
	}

	log.Debug(hotelsDto)
	log.Debug("HOLA ACA")

	c.JSON(http.StatusOK, hotelsDto)
}

//654513064fd3abb838f8c48d
