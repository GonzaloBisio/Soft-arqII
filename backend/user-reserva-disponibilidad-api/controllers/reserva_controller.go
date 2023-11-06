package controllers

import (
	"log"
	"net/http"
	"strconv"
	dtos "user-reserva-disponibilidad-api/dtos/reserva_dto"
	"user-reserva-disponibilidad-api/services"
	se "user-reserva-disponibilidad-api/services"

	"github.com/gin-gonic/gin"
)

type CustomError struct {
	Message string
	Code    int
}

func (e CustomError) Error() string {
	return e.Message
}

func (e CustomError) Status() int {
	return e.Code
}

func NewReserva(c *gin.Context) {
	var newReserva dtos.ReservaDto
	log.Println("llegue al controller")

	if err := c.ShouldBindJSON(&newReserva); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Println("pasa al service")
	log.Println(newReserva)
	newReserva, err := services.ReservationService.NewReserva(newReserva)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.AddParam("id", strconv.Itoa(newReserva.Id))
	c.JSON(http.StatusOK, newReserva)
}

func CheckStatus(ctx *gin.Context) {
	log.Println("CheckStatus")
	ctx.JSON(http.StatusOK, `OK`)

}

func GetReservaById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parámetro inválido: ID no es un número entero"})
		return
	}

	create := se.ReservationService.GetReservaById(id)

	ctx.JSON(http.StatusOK, create)
}

func GetReservas(ctx *gin.Context) {
	reservasDto, err := se.ReservationService.GetReservas()

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, reservasDto)
}

func Disponibilidad_de_reserva(ctx *gin.Context) {
	var create dtos.ReservaDto

	idH := ctx.Param("idHotel")
	inicio := ctx.Param("inicio")
	final := ctx.Param("final")

	idU, _ := strconv.Atoi(ctx.Param("idUser"))

	create.HotelId = idH
	create.StartDate = inicio
	create.EndDate = final
	create.UserId = idU

	/*reservationDTO, err := se.ReservationService.Disponibilidad_de_reserva(create)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, reservationDTO)
	}
	*/
}

func GetReservasByUserId(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("user_id"))

	reservasDto, err := se.ReservationService.GetReservasByUserId(id)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, reservasDto)
}
