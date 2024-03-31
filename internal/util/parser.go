package util

import (
	"camarinb2096/unit-testing-course/internal/dto"
	"encoding/json"
	"fmt"
)

func ParserData(data []byte) interface{} {
	fmt.Println("Parsing data...")

	var pokeList dto.PokeListDto

	json.Unmarshal(data, &pokeList)

	fmt.Println("Data parsed:", pokeList)

	return data
}
