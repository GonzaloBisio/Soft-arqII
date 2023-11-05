package services

import (
	"encoding/json"
	"fmt"
	"io"

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

	_, err := http.Get(fmt.Sprintf("%shotelId/%s", config.HotelUrl, id)) //Link a la api de hotel_list
	if err == nil {
		return fmt.Errorf("Hotel aun existente")
	}

	err = client.DeleteFromId(id)

	if err != nil {
		return err
	}

	return nil

}

//HOla
