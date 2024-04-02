package dto

type Response struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Count   int       `json:"count"`
	Data    []Results `json:"data"`
}
