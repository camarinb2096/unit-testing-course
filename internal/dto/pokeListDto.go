package dto

type PokeListDto struct {
	Name    string      `json:"name"`
	Url     string      `json:"url"`
	Results interface{} `json:"results"`
}
