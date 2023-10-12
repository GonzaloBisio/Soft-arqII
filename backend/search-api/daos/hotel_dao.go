package daos

import (
	"fmt"
	"search-api/models"

	solr "github.com/rtt/Go-Solr"
)

type HotelDao interface{
    Get(id string) (*models.Hotel, error)
    Create(hotel *models.Hotel) error
    Update(hotel *models.Hotel) error
    GetAll() ([]*models.Hotel, error)
    GetByCity(city string) ([]*models.Hotel, error)
}

type HotelSolrDao struct{}

func NewHotelSolrDAO() HotelDao {
	return &HotelSolrDao{}
}

func (dao *HotelSolrDao) Get(id string) (*models.Hotel, error) {
    q := &solr.Query{
        Params: solr.URLParamMap{
            "q":    []string{fmt.Sprintf("id:%s", id)},
            "rows": []string{"1"},
        },
    }

    res, err := db.SolrClient.Search(q)
    if err != nil {
        return nil, err
    }

    if len(res.Results.Docs) > 0 {
        doc := res.Results.Docs[0]
        hotel := &models.Hotel{
            ID:          doc["id"].(string),
            Name:        doc["name"].(string),
            Description: doc["description"].(string),
        }
        return hotel, nil
    }

    return nil, fmt.Errorf("Hotel not found")
}


func (dao *HotelSolrDao) Create(hotel *models.Hotel) error {
	// Crear un mapa que representa el documento del hotel
	hotelDocument := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{
				"id":          hotel.ID,
				"name":        hotel.Name,
				"description": hotel.Description,
			},
		},
	}

	// Inserta el nuevo documento en Solr
	_, err := db.SolrClient.Update(hotelDocument, true) // El segundo parámetro "true" realiza una confirmación inmediata
	if err != nil {
		return err
	}
	return nil
}

func (dao *HotelSolrDao) Update(hotel *models.Hotel) error {
	// Crear un mapa que representa los cambios en el documento del hotel
	hotelDocument := map[string]interface{}{
		"update": []interface{}{
			map[string]interface{}{
				"id":          hotel.ID,
				"name":        hotel.Name,
				"description": hotel.Description,
			},
		},
	}

	// Actualiza el documento del hotel en Solr
	_, err := db.SolrClient.Update(hotelDocument, true) 
	if err != nil {
		return err
	}
	return nil
}

func (dao *HotelSolrDao) GetAll() ([]*models.Hotel, error) {
    q := &solr.Query{
        Params: solr.URLParamMap{
            "q": []string{"*:*"},
            "rows": []string{"100"},
        },
    }

    res, err := db.SolrClient.Search(q)
    if err != nil {
        return nil, err
    }

    var hotels []*models.Hotel
    for _, doc := range res.Results.Docs {
        hotel := &models.Hotel{
            ID:          doc["id"].(string),
            Name:        doc["name"].(string),
            Description: doc["description"].(string),
        }
        hotels = append(hotels, hotel)
    }

    return hotels, nil
}

func (dao *HotelSolrDao) GetByCity(city string) ([]*models.Hotel, error) {
    q := &solr.Query{
        Params: solr.URLParamMap{
            "q":   []string{fmt.Sprintf("city:%s", city)},
            "rows": []string{"100"},
        },
    }

    res, err := db.SolrClient.Search(q)
    if err != nil {
        return nil, err
    }

    var hotels []*models.Hotel
    for _, doc := range res.Results.Docs {
        hotel := &models.Hotel{
            ID:          doc["id"].(string),
            Name:        doc["name"].(string),
            Description: doc["description"].(string),
        }
        hotels = append(hotels, hotel)
    }

    return hotels, nil
}

func (dao *HotelSolrDao) GetByAvailability(city string, checkin string, checkout string) ([]*models.Hotel, error) {
    q := &solr.Query{
        Params: solr.URLParamMap{
            "q":   []string{fmt.Sprintf("city:%s AND availability:true", city)},
            "rows": []string{"100"},
        },
    }

    res, err := db.SolrClient.Search(q)
    if err != nil {
        return nil, err
    }

    var hotels []*models.Hotel
    for _, doc := range res.Results.Docs {
        hotel := &models.Hotel{
            ID:          doc["id"].(string),
            Name:        doc["name"].(string),
            Description: doc["description"].(string),
        }
        hotels = append(hotels, hotel)
    }

    return hotels, nil
}
