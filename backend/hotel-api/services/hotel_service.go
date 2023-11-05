package services

import (
	"fmt"
	dao "hotel-api/dao"
	"hotel-api/dto"
	"hotel-api/models"
	"hotel-api/utils/errors"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type hotelService struct {
	Collection *mongo.Collection
}

type hotelServiceInterface interface {
	GetHotels() (dto.HotelsDto, errors.ApiError)
	GetHotelById(string) (dto.HotelDTO, errors.ApiError)
	InsertHotel(dto.HotelDTO) (dto.HotelDTO, errors.ApiError)
	UpdateHotel(dto.HotelDTO) (dto.HotelDTO, errors.ApiError)
	DeleteHotelById(string) errors.ApiError
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotels() (dto.HotelsDto, errors.ApiError) {
	hotels, err := dao.GetAll()
	if err != nil {
		return dto.HotelsDto{}, errors.NewInternalServerApiError("Ningun hotel encontrado", err)
	}

	var hotelDtos = make([]dto.HotelDTO, 0)
	for _, hotel := range hotels {
		hotelDto := models.Hotel{
			ID:          hotel.ID,
			Name:        hotel.Name,
			Description: hotel.Description,
		}
		hotelDtos = append(hotelDtos, ConvertToHotelDTO(hotelDto))
	}

	final := dto.HotelsDto{
		Hotels: hotelDtos,
	}

	return final, nil
}

func (s *hotelService) GetHotelById(id string) (dto.HotelDTO, errors.ApiError) {
	var hotel dto.HotelDTO
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return hotel, errors.NewBadRequestApiError(err.Error())
	}

	mhotel, err := dao.GetHotelById(ID) // Convierte ObjectID a cadena

	if err != nil {
		return hotel, errors.NewBadRequestApiError(err.Error())
	}
	hotel = ConvertToHotelDTO(mhotel)
	return hotel, nil
}

func (s *hotelService) InsertHotel(hotel dto.HotelDTO) (dto.HotelDTO, errors.ApiError) {

	hotelInsertado, err := dao.Insert(ConvertToHotelModel(hotel))

	if err != nil {
		return hotel, errors.NewInternalServerApiError("Error al insertar el hotel en la base de datos", err)
	}
	return ConvertToHotelDTO(hotelInsertado), nil
}

func (s *hotelService) UpdateHotel(hotel dto.HotelDTO) (dto.HotelDTO, errors.ApiError) {

	_, err := dao.Update(ConvertToHotelModel(hotel))
	if err != nil {
		return dto.HotelDTO{}, errors.NewInternalServerApiError("Error al actualizar el hotel en la base de datos", err)
	}
	return hotel, nil
}

func (s *hotelService) DeleteHotelById(id string) errors.ApiError {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.NewBadRequestApiError(err.Error())
	}

	err = dao.DeleteHotelById(objectID) // Convierte ObjectID a cadena

	if err != nil {
		return errors.NewBadRequestApiError(err.Error())
	}

	return nil
}

func handleFileupload(c *fiber.Ctx) error {

	// parse incomming image file

	file, err := c.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	// generate new uuid for image name
	uniqueId := uuid.New()

	// remove "- from imageName"

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename

	fileExt := strings.Split(file.Filename, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// save image to ./images dir
	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// generate image url to serve to client using CDN

	imageUrl := fmt.Sprintf("http://localhost:4000/images/%s", image)

	// create meta data and send to client

	data := map[string]interface{}{

		"imageName": image,
		"imageUrl":  imageUrl,
		"header":    file.Header,
		"size":      file.Size,
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}

func handleDeleteImage(c *fiber.Ctx) error {
	// extract image name from params
	imageName := c.Params("imageName")

	// delete image from ./images
	err := os.Remove(fmt.Sprintf("./images/%s", imageName))
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server Error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image deleted successfully", "data": nil})
}

func ConvertToHotelDTO(hotel models.Hotel) dto.HotelDTO {
	return dto.HotelDTO{
		ID:          hotel.ID.Hex(),
		Name:        hotel.Name,
		Description: hotel.Description,
		Fotos:       hotel.Fotos,
		Amenities:   hotel.Amenities,
	}
}

func ConvertToHotelModel(hotelDTO dto.HotelDTO) models.Hotel {
	id, _ := primitive.ObjectIDFromHex(hotelDTO.ID)

	return models.Hotel{
		ID:          id,
		Name:        hotelDTO.Name,
		Description: hotelDTO.Description,
		Fotos:       hotelDTO.Fotos,
		Amenities:   hotelDTO.Amenities,
	}
}
