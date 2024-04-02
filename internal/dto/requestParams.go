package dto

type RequestParams struct {
	Id    string `json:"id"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}
