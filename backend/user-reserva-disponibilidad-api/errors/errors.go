package errors

import (
	"fmt"
	"net/http"

	json "github.com/json-iterator/go"
)

type CauseList []interface{}

type ErrorApi interface {
	Message() string
	Code() string
	Status() int
	Cause() CauseList
	Error() string
}

type ErrApi struct {
	ErrorMessage string    `json:"message"`
	ErrorCode    string    `json:"error"`
	ErrorStatus  int       `json:"status"`
	ErrorCause   CauseList `json:"cause"`
}

func (c CauseList) ToString() string { //pasar la lista de causas a una cadena
	return fmt.Sprint(c)
}

func (e ErrApi) Code() string { //implementa el metodo Code de la interfaz ErrorApi
	return e.ErrorCode
}

func (e ErrApi) Error() string { //implementa el metodo Error de la interfaz ErrorApi
	return fmt.Sprintf("Message: %s;Error Code: %s;Status: %d;Cause: %v", e.ErrorMessage, e.ErrorCode, e.ErrorStatus, e.ErrorCause)
}

func (e ErrApi) Status() int { //implementa el metodo Status de la interfaz ErrorApi
	return e.ErrorStatus
}

func (e ErrApi) Cause() CauseList { //implementa el metodo Cause de la interfaz ErrorApi
	return e.ErrorCause
}

func (e ErrApi) Message() string { //implementa el metodo Message de la interfaz ErrorApi
	return e.ErrorMessage
}

//Ahora definimos funciones para crear diferentes tipos de errores API

func NewErrorApi(message string, error string, status int, cause CauseList) ErrorApi { //crea un nuevo error con mensaje, codigo, estado y causas
	return ErrApi{message, error, status, cause}
}

func NewNotFoundErrorApi(message string) ErrorApi { //crear un nuevo error API de un recurso no encontrado
	return ErrApi{message, "not_found", http.StatusNotFound, CauseList{}}
}

func NewTooManyRequestsError(message string) ErrorApi { //crear un nuevo error API de demasiadas solicitudes
	return ErrApi{message, "too_many_requests", http.StatusTooManyRequests, CauseList{}}
}

func NewBadRequestErrorApi(message string) ErrorApi { //crear un nuevo error API de una solicitud incorrecta
	return ErrApi{message, "bad_request", http.StatusBadRequest, CauseList{}}
}

func NewValidationErrorApi(message string, error string, cause CauseList) ErrorApi { //crear un nuevo error API de error de validacion
	return ErrApi{message, error, http.StatusBadRequest, cause}
}

func NewMethodNotAllowedErrorApi() ErrorApi { //crear un nuevo error API de un metodo no permitido
	return ErrApi{"Method not allowed", "method_not_allowed", http.StatusMethodNotAllowed, CauseList{}}
}

func NewInternalServerErrorApi(message string, err error) ErrorApi { //crear un nuevo error API de error interno del servidor
	cause := CauseList{}
	if err != nil {
		cause = append(cause, err.Error())
	}
	return ErrApi{message, "internal_server_error", http.StatusInternalServerError, cause}
}

func NewForbiddenErrorApi(message string) ErrorApi { //crear un nuevo error API de acceso denegado
	return ErrApi{message, "forbidden", http.StatusForbidden, CauseList{}}
}

func NewUnauthorizedErrorApi(message string) ErrorApi { //crear un nuevo error API de autenticacion no autorizada
	return ErrApi{message, "unauthorized_scopes", http.StatusUnauthorized, CauseList{}}
}

func NewConflictErrorApi(id string) ErrorApi { //crear un nuevo error API de de conflicto para un ID especifico
	return ErrApi{"Can't update " + id + " due to a conflict error", "conflict_error", http.StatusConflict, CauseList{}}
}

func NewErrorApiFromBytes(data []byte) (ErrorApi, error) { //Toma un arreglo de bytes que representa un error API en formato JSON y lo convierte en
	//una instancia de ErrorApi mediante la decodificacion JSON
	err := ErrApi{}
	e := json.Unmarshal(data, &err)
	return err, e
}
