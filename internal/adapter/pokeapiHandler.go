package adapter

import (
	"camarinb2096/unit-testing-course/internal/dto"
	"camarinb2096/unit-testing-course/internal/util"
	"fmt"
	"io"
	"net/http"
)

func GetPokemonList(requestParams dto.RequestParams) dto.Response {

	uri := fmt.Sprintf("https://pokeapi.co/api/v2/%s", requestParams.Id)

	if requestParams.Limit > 0 && requestParams.Page > 0 {
		uri = fmt.Sprintf("%s?limit=%d&offset=%d", uri, requestParams.Limit, (requestParams.Page-1)*requestParams.Limit)
	}

	response, err := http.Get(uri)
	if err != nil {
		return dto.Response{
			Status:  response.StatusCode,
			Message: err.Error(),
		}
	}

	if response.StatusCode != http.StatusOK {
		return dto.Response{
			Status: response.StatusCode,
		}
	}

	defer response.Body.Close()

	bufferData, err := io.ReadAll(response.Body)
	if err != nil {
		return dto.Response{
			Status:  500,
			Message: "error reading pokemon data",
		}
	}

	data := util.ParserData(bufferData)

	return dto.Response{
		Status:  response.StatusCode,
		Message: "Pokemon data retrieved successfully!",
		Count:   data.(dto.PokeList).Count,
		Data:    data.(dto.PokeList).Results,
	}
}
