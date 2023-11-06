package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"net/http"
	client "search-api/Client"
	"search-api/config"

	"search-api/dto"

	"github.com/stevenferrer/solr-go"
)

type solrService struct {
	SolrClient *solr.JSONClient
	Collection string
}

type solrServiceInterface interface {
	AddFromId(id string) error
	DeleteFromId(id string) error
	GetQuery(query string, field string) (dto.HotelsDto, error)
}

var (
	SolrService solrServiceInterface
)

func init() {
	SolrService = &solrService{}
}

func (s *solrService) AddFromId(id string) error {
	var hotelDto dto.HotelDTO

	resp, err := http.Get(fmt.Sprintf("%shotelId/%s", config.HotelUrl, id)) //Link a la api de hotel_list
	if err != nil {
		return err
	}
	var body []byte

	body, _ = io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &hotelDto)

	if err != nil {

		return err
	}

	hotelDto, err = client.AddFromId(hotelDto)

	if err != nil {
		return err
	}

	return nil

}

func (s *solrService) DeleteFromId(id string) error {

	err := client.DeleteFromId(id)

	if err != nil {
		return err
	}

	return nil

}

func (s *solrService) GetQuery(query string, field string) (dto.HotelsDto, error) {
	var response dto.SolrResponseDto
	var hotelsDto dto.HotelsDto
	q, err := http.Get(fmt.Sprintf("http://%s/solr/hotelSearch/select?q=%s%s%s", config.SolrURL, field, "%3A", query))

	if err != nil {
		return hotelsDto, err
	}

	defer q.Body.Close()
	err = json.NewDecoder(q.Body).Decode(&response)
	if err != nil {
		log.Printf("Response Body: %s", q.Body) // Add this line
		log.Printf("Error: %s", err.Error())
		return hotelsDto, err
	}
	hotelsDto = response.Response.Docs

	log.Printf("hotels:", hotelsDto)

	return hotelsDto, nil
}
