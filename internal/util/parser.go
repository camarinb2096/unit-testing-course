package util

import (
	"camarinb2096/unit-testing-course/internal/dto"
	"encoding/json"
)

func ParserData(data []byte) interface{} {
	var pokeList dto.PokeList
	json.Unmarshal(data, &pokeList)

	return pokeList
}
