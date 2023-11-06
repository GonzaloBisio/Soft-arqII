package services

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	client "search-api/Client"
	"search-api/config"
	"search-api/dto"
	"strconv"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/stevenferrer/solr-go"
)

type solrService struct {
	SolrClient *solr.JSONClient
	Collection string
}

type solrServiceInterface interface {
	AddFromId(id string) error
	DeleteFromId(id string) error
	SerchQuery(query string) (dto.HotelsDto, error)
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

func (s *solrService) SerchQuery(query string) (dto.HotelsDto, error) {
	var hotelsDto dto.HotelsDto
	queryParams := strings.Split(query, "_")

	numParams := len(queryParams)

	log.Printf("Params: %d", numParams)

	field, query := queryParams[0], queryParams[1]

	log.Printf("%s and %s", field, query)

	hotelsDto, err := client.SerchQuery(query, field)
	if err != nil {
		return hotelsDto, err
	}

	if numParams == 4 {

		startdateQuery, enddateQuery := queryParams[2], queryParams[3]
		startdateSplit := strings.Split(startdateQuery, "-")
		enddateSplit := strings.Split(enddateQuery, "-")
		startdate := fmt.Sprintf("%s%s%s", startdateSplit[0], startdateSplit[1], startdateSplit[2])
		enddate := fmt.Sprintf("%s%s%s", enddateSplit[0], enddateSplit[1], enddateSplit[2])

		sDate, _ := strconv.Atoi(startdate)
		eDate, _ := strconv.Atoi(enddate)

		// Create a channel to collect results
		resultsChan := make(chan dto.HotelDto, len(hotelsDto))

		// Create a WaitGroup
		var wg sync.WaitGroup
		var hotel dto.HotelDto

		// Iterate through each hotel and make concurrent API calls
		for _, hotel = range hotelsDto {
			wg.Add(1) // Increment the WaitGroup counter for each Goroutine
			go func(hotel dto.HotelDto) {
				defer wg.Done() // Decrement the WaitGroup counter when Goroutine is done

				// Make API call for each hotel and send the hotel ID
				result, err := GetHotelInfo(hotel.Id, sDate, eDate) // Assuming you have a function to get hotel info
				if err != nil {
					result = false
				}

				var response dto.HotelDto

				if result == true {
					response = hotel
				}

				resultsChan <- response
			}(hotel)
		}

		// Create a slice to store the results
		var hotelResults dto.HotelsDto

		// Start a Goroutine to close the channel when all Goroutines are done
		go func() {
			wg.Wait()          // Wait for all Goroutines to finish
			close(resultsChan) // Close the channel when all Goroutines are done
		}()

		// Collect results from the channel
		for response := range resultsChan {
			hotelResults = append(hotelResults, response)
		}

		return hotelResults, nil

	}

	return hotelsDto, nil
}

func GetHotelInfo(id string, startdate int, enddate int) (bool, error) {

	resp, err := http.Get(fmt.Sprintf("http://%s/hotel/availability/%s/%d/%d", config.ReservationlUrl, id, startdate, enddate))

	if err != nil {
		return false, err
	}

	var body []byte
	body, _ = io.ReadAll(resp.Body)

	var responseDto dto.AvailabilityResponse
	err = json.Unmarshal(body, &responseDto)

	if err != nil {
		log.Debugf("error in unmarshal")
		return false, err
	}

	status := responseDto.Status
	return status, nil
}
