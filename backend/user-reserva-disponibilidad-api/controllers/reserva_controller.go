package controllers

import (
	"net/http"
	"strconv"
	"user-reserva-disponibilidad-api/dtos/reserva_dto"
	se "user-reserva-disponibilidad-api/services"

	"github.com/gin-gonic/gin"
)

func NewReserva(ctx *gin.Context) {
	var newReserva dtos.ReservaDto

	if err := ctx.ShouldBindJSON(&newReserva); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar la solicitud"})
		return
	}

	reservationDTO, err := se.ReservationService.NewReserva(newReserva)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, reservationDTO)
	ctx.AddParam("id", strconv.Itoa(newReserva.Id))
	ctx.JSON(http.StatusOK, newReserva)
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

	reservationDTO, err := se.ReservationService.Disponibilidad_de_reserva(create)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, reservationDTO)
	}
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
