package services

import (
	"encoding/json"
	"io"
	"net/http"
	client "search-api/Client"

	"search-api/dto"

	"github.com/stevenferrer/solr-go"
)

type solrService struct {
	SolrClient *solr.JSONClient
	Collection string
}

type solrServiceInterface interface {
	AddFromId(id string) error
	Delete(id string) error
}

var (
	SolrService solrServiceInterface
)

func (*solrService) AddFromId(id string) error {
	var hotelDto dto.HotelDto
	resp, err := http.Get("") //Link a la api de hotel_list

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

//HOla
