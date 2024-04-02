package adapter

import (
	"camarinb2096/unit-testing-course/internal/dto"
	"camarinb2096/unit-testing-course/internal/util"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestGetPokemonList(t *testing.T) {
	t.Run("GetPokemonList", func(t *testing.T) {
		// We expect the response to be a list of pokemons with a length of 10 elements
		got := GetPokemonList(dto.RequestParams{Id: "berry", Page: 1, Limit: 10})
		if len(got.Data) != 10 {
			t.Errorf("Expected 10 pokemons, got %d", len(got.Data))
		}

		got = GetPokemonList(dto.RequestParams{Id: "berry", Page: 1, Limit: 10})
		if got.Status != 200 {
			t.Errorf("Expected 200, got %d", got.Status)
		}

	})
}

func TestGetPokemonListSuccesMock(t *testing.T) {
	c := require.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	body, err := os.ReadFile("sample/poke_api_list.json")
	c.NoError(err)

	httpmock.RegisterResponder("GET", "https://pokeapi.co/api/v2/berry?limit=10&offset=0", httpmock.NewBytesResponder(200, body))

	got := GetPokemonList(dto.RequestParams{Id: "berry", Page: 1, Limit: 10})
	c.NoError(err)

	var expected dto.Response

	response := util.ParserData([]byte(body))
	c.NoError(err)

	expected = dto.Response{
		Status:  200,
		Message: "Pokemon data retrieved successfully!",
		Count:   response.(dto.PokeList).Count,
		Data:    response.(dto.PokeList).Results,
	}

	c.Equal(expected, got)
}

func TestGetPokemonListInternalServerError(t *testing.T) {
	c := require.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://pokeapi.co/api/v2/berry?limit=10&offset=0", httpmock.NewStringResponder(500, ""))

	got := GetPokemonList(dto.RequestParams{Id: "berry", Page: 1, Limit: 10})
	c.Equal(500, got.Status)
}
