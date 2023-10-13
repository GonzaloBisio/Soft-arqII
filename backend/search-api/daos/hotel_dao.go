package daos

import (
	"fmt"
	db "search-api/db"
	model "search-api/models"

	solr "github.com/rtt/Go-Solr"
)

type HotelDao interface {
    GetById(id string) (*model.Hotel, error)
    CreateHotel(hotel *model.Hotel) error
    UpdateHotel(hotel *model.Hotel) error
    GetAllHotels() ([]*model.Hotel, error)
    GetByCity(city string) ([]*model.Hotel, error)
}

type HotelSolrDao struct{}

func NewHotelSolrDAO() HotelDao {
    return &HotelSolrDao{}
}

func (dao *HotelSolrDao) GetById(id string) (*model.Hotel, error) {
    q := &solr.Query{
        Params: solr.URLParamMap{
            "q":    []string{fmt.Sprintf("id:%s", id)},
            "rows": []string{"1"},
        },
    }

    res, err := db.SolrClient.Select(q)
    if err != nil {
        return nil, err
    }

    if len(res.Results.Collection) > 0 {
		doc := res.Results.Collection[0]
		hotel := &model.Hotel{
			ID:          doc.Fields["id"].(string),
			Name:        doc.Field("name").([]interface{})[0].(string),
			City:        doc.Field("city").([]interface{})[0].(string),
			Description: doc.Field("description").([]interface{})[0].(string),
			Thumbnail:   doc.Field("thumbnail").([]interface{})[0].(string),
			Images:      getStringsFromInterface(doc.Field("images")),
			Amenities:   getStringsFromInterface(doc.Field("amenities")),
		}
		return hotel, nil
	}

    return nil, fmt.Errorf("Hotel not found")
}

func (dao *HotelSolrDao) CreateHotel(hotel *model.Hotel) error {
    // documento que representa el hotel
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


func (dao *HotelSolrDao) UpdateHotel(hotel *model.Hotel) error {
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


func (dao *HotelSolrDao) GetAllHotels() ([]*model.Hotel, error) {
    q := &solr.Query{
        Params: solr.URLParamMap{
            "q":   []string{"*:*"},
            "rows": []string{"100"},
        },
    }

    res, err := db.SolrClient.Select(q)
    if err != nil {
        return nil, err
    }

    var hotels []*model.Hotel
    for _, doc := range res.Results.Collection {
        hotel := &model.Hotel{
            ID:          doc.Fields["id"].(string),
            Name:        doc.Field("name").([]interface{})[0].(string),
            City:        doc.Field("city").([]interface{})[0].(string),
            Description: doc.Field("description").([]interface{})[0].(string),
            Thumbnail:   doc.Field("thumbnail").([]interface{})[0].(string),
            Images:      getStringsFromInterface(doc.Field("images")),
            Amenities:   getStringsFromInterface(doc.Field("amenities")),
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

    res, err := db.SolrClient.Select(q)
    if err != nil {
        return nil, err
    }

    var hotels []*model.Hotel
    for _, doc := range res.Results.Collection {
        hotel := &model.Hotel{
            ID:          doc.Fields["id"].(string),
            Name:        doc.Field("name").([]interface{})[0].(string),
            City:        doc.Field("city").([]interface{})[0].(string),
            Description: doc.Field("description").([]interface{})[0].(string),
            Thumbnail:   doc.Field("thumbnail").([]interface{})[0].(string),
            Images:      getStringsFromInterface(doc.Field("images")),
            Amenities:   getStringsFromInterface(doc.Field("amenities")),
        }
        hotels = append(hotels, hotel)
    }

    return hotels, nil
}


//quedaria hacer segun disponibilidad 

//esto sirve para usar en la funciones de arriba, para tratar algunos campos como listas,
//entonces se toma el primer elemento digamos para luego obtener todo el valor de la cadena completa
func getStringsFromInterface(data interface{}) []string {
    if dataSlice, ok := data.([]interface{}); ok {
        result := make([]string, len(dataSlice))
        for i, item := range dataSlice {
            if str, isString := item.(string); isString {
                result[i] = str
            }
        }
        return result
    }
    return []string{}
}
