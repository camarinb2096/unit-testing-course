package adapter

import "testing"

func TestGetPokemon(t *testing.T) {
	t.Run("GetPokemon", func(t *testing.T) {
		// GetPokemon is a function that returns an error
		err := GetPokemon("pokemon/1")
		// We expect the error to be nil
		if err != nil {
			t.Errorf("GetPokemon() error = %v; want nil", err)
		}
	})

	t.Run("GetPokemon", func(t *testing.T) {
		err := GetPokemon("pokemon/go")
		// We expect the error to be not nil
		if err == nil {
			t.Errorf("GetPokemon() error = nil; want not nil")
		}
	})

}
