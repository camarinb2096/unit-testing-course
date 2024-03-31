package adapter

import (
	"camarinb2096/unit-testing-course/internal/util"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetPokemon(id string) error {
	log.Println("Getting pokemon data...")

	uri := fmt.Sprintf("https://pokeapi.co/api/v2/%s", id)

	response, err := http.Get(uri)
	if err != nil {
		log.Println("Error:", err)
	}

	defer response.Body.Close()

	bufferData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error:", err)
	}

	util.ParserData(bufferData)

	return nil

}
