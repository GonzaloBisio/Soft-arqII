package daos

import (
	"fmt"
	model "search-api/models"
	"search-api/utils/db"

	solr "github.com/rtt/Go-Solr"
)

type HotelDao interface {
    Get(id string) (*model.Hotel, error)
    Create(hotel *model.Hotel) error
    Update(hotel *model.Hotel) error
    GetAll() ([]*model.Hotel, error)
    GetByCity(city string) ([]*model.Hotel, error)
}

type HotelSolrDao struct{}

func NewHotelSolrDAO() HotelDao {
    return &HotelSolrDao{}
}

func (dao *HotelSolrDao) Get(id string) (*model.Hotel, error) {
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
        hotel := &model.Hotel{
            ID:          doc["id"].(string),
            Name:        doc["name"].(string),
            City:        doc["city"].(string),
            Description: doc["description"].(string),
            Thumbnail:   doc["thumbnail"].(string),
            Images:      doc["images"].([]string),
            Amenities:   doc["amenities"].([]string),
        }
        return hotel, nil
    }

    return nil, fmt.Errorf("Hotel not found")
}

func (dao *HotelSolrDao) Create(hotel *model.Hotel) error {
    // Crear un mapa que representa el documento del hotel
    hotelDocument := map[string]interface{}{
        "add": []interface{}{
            map[string]interface{}{
                "id":          hotel.ID,
                "name":        hotel.Name,
                "city":        hotel.City,
                "description": hotel.Description,
                "thumbnail":   hotel.Thumbnail,
                "images":      hotel.Images,
                "amenities":   hotel.Amenities,
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


func (dao *HotelSolrDao) Update(hotel *model.Hotel) error {
    // Crear un mapa que representa los cambios en el documento del hotel
    hotelDocument := map[string]interface{}{
        "update": []interface{}{
            map[string]interface{}{
                "id":          hotel.ID,
                "name":        hotel.Name,
                "city":        hotel.City,
                "description": hotel.Description,
                "thumbnail":   hotel.Thumbnail,
                "images":      hotel.Images,
                "amenities":   hotel.Amenities,
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


func (dao *HotelSolrDao) GetAll() ([]*model.Hotel, error) {
    q := &solr.Query{
        Params: solr.URLParamMap{
            "q":   []string{"*:*"},
            "rows": []string{"100"},
        },
    }

    res, err := db.SolrClient.Search(q)
    if err != nil {
        return nil, err
    }

    var hotels []*model.Hotel
    for _, doc := range res.Results.Docs {
        hotel := &model.Hotel{
            ID:          doc["id"].(string),
            Name:        doc["name"].(string),
            City:        doc["city"].(string),
            Description: doc["description"].(string),
            Thumbnail:   doc["thumbnail"].(string),
            Images:      doc["images"].([]string),
            Amenities:   doc["amenities"].([]string),
        }
        hotels = append(hotels, hotel)
    }

    return hotels, nil
}

func (dao *HotelSolrDao) GetByCity(city string) ([]*model.Hotel, error) {
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

    var hotels []*model.Hotel
    for _, doc := range res.Results.Docs {
        hotel := &model.Hotel{
            ID:          doc["id"].(string),
            Name:        doc["name"].(string),
            City:        doc["city"].(string),
            Description: doc["description"].(string),
            Thumbnail:   doc["thumbnail"].(string),
            Images:      doc["images"].([]string),
            Amenities:   doc["amenities"].([]string),
        }
        hotels = append(hotels, hotel)
    }

    return hotels, nil
}


//quedaria hacer segun disponibilidad 