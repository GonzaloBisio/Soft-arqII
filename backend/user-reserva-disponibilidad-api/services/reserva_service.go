package services

import (
	"time"
	cl "user-reserva-disponibilidad-api/clients"
	reservationDTO "user-reserva-disponibilidad-api/dtos/reserva_dto"
	e "user-reserva-disponibilidad-api/errors"
	"user-reserva-disponibilidad-api/model"
)

type reservationService struct {
}

type reservationServicesInterface interface {
	NewReserva(dto reservationDTO.ReservaDto) (reservationDTO.ReservaDto, error)
	InsertReserva(dto reservationDTO.ReservaDto) (reservationDTO.ReservaDto, error)
	GetReservaById(int) reservationDTO.ReservaDto
	GetReservas() ([]reservationDTO.ReservaDto, e.ErrorApi)
	GetReservasByUserId(int) ([]reservationDTO.ReservaDto, e.ErrorApi)
	Disponibilidad_de_reserva(dto reservationDTO.ReservaDto) (bool, error)
}

var (
	ReservationService reservationServicesInterface
	Layoutd            = "2006-01-02"
)

func (s *reservationService) NewReserva(reserva reservationDTO.ReservationCreateDto) (reservationDTO.ReservaDto, error) {
	var Mreserva model.Reservation
	var rf reservationDTO.ReservaDto
	Mreserva.HotelID = reserva.HotelId

	parseInitial, err := time.Parse(Layoutd, reserva.InitialDate)
	if err != nil {
		return rf, err
	}
	Mreserva.InitialDate = parseInitial

	parseFinal, err := time.Parse(Layoutd, reserva.FinalDate)
	if err != nil {
		return rf, err
	}
	Mreserva.FinalDate = parseFinal

	Mreserva.UserID = reserva.UserId

	if parseFinal.Before(parseInitial) {
		return rf, e.NewBadRequestErrorApi("Fecha inicial antes de la final")
	}

	rf.EndDate = Mreserva.FinalDate.Format(Layoutd)
	rf.HotelId = Mreserva.HotelID
	rf.StartDate = Mreserva.InitialDate.Format(Layoutd)
	rf.UserId = Mreserva.UserID

	return rf, nil
}

func (s *reservationService) GetReservaById(id int) reservationDTO.ReservaDto {
	var re reservationDTO.ReservaDto
	re.Id = id
	c := cl.GetReservaById(id)
	re.EndDate = c.FinalDate.Format(Layoutd)
	re.HotelId = c.HotelID
	re.Id = c.Id
	re.StartDate = c.InitialDate.Format(Layoutd)
	re.UserId = c.UserID
	return re
}

func (s *reservationService) GetReservas() ([]reservationDTO.ReservaDto, e.ErrorApi) {
	var reservas model.Reservations = cl.GetReservas()
	reservasList := make([]reservationDTO.ReservaDto, 0)
	for _, reserva := range reservas {
		var reservaDto reservationDTO.ReservaDto
		id := reserva.Id
		reservaDto = s.GetReservaById(id)
		reservasList = append(reservasList, reservaDto)
	}
	return reservasList, nil
}

func (s *reservationService) Disponibilidad_de_reserva(reserva reservationDTO.ReservationCreateDto) (bool, error) {
	var Mreserva model.Reservation
	Mreserva.HotelID = reserva.HotelId

	parseInitial, err := time.Parse(Layoutd, reserva.InitialDate)
	if err != nil {
		return false, err
	}
	Mreserva.InitialDate = parseInitial

	parseFinal, err := time.Parse(Layoutd, reserva.FinalDate)
	if err != nil {
		return false, err
	}
	Mreserva.FinalDate = parseFinal

	Mreserva.UserID = reserva.UserId

	if parseFinal.Before(parseInitial) {
		return false, e.NewBadRequestErrorApi("Fecha inicial antes de la final")
	}

	if cl.ComprobarReserva(Mreserva) {
		return true, nil
	} else {
		return false, e.NewBadRequestErrorApi("Las fechas no están disponibles") //Completar este error adecuadamente
	}
}

func (s *reservationService) GetReservasByUserId(id int) ([]reservationDTO.ReservaDto, e.ErrorApi) {
	var reservas model.Reservations = cl.GetReservasByUserId(id)
	reservasList := make([]reservationDTO.ReservaDto, 0)
	for _, reserva := range reservas {
		var reservaDto reservationDTO.ReservaDto
		id := reserva.Id
		reservaDto = s.GetReservaById(id)
		//log.Println(reservaDto.Habitacion)
		reservasList = append(reservasList, reservaDto)
	}
	return reservasList, nil
}
