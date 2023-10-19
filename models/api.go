package models

type NotFoundMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
