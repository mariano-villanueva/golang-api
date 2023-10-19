package utils

import (
	"api/models"
	"net/http"
)

func NotFoundMessage() models.NotFoundMessage {
	var mensaje = models.NotFoundMessage{
		Status:  http.StatusNotFound,
		Message: "El recurso solicitado no ha sido encontrado.",
	}
	return mensaje
}
