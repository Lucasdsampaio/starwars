package api

import (
	"testing"
)

func TestGetPlanetsFromResponse(t *testing.T) {

	responseAPI, err := AsyncGet()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	planets, err := GetPlanetsFromResponse(responseAPI)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if planets == nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(planets) != 60 {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestCreatePlanetDB(t *testing.T) {
	inputInPlanets := CreatePlanetInput{
		Name:    "Tatooine",
		Climate: "arid",
		Terrain: "desert",
	}
	inputWithoutPlanets := CreatePlanetInput{
		Name:    "Earth",
		Climate: "temperate",
		Terrain: "oceans, reefs, islands",
	}
	planets := []Planet{Planet{
		Name:    "Tatooine",
		Climate: "arid",
		Terrain: "desert",
		Films: []string{
			"http://swapi.dev/api/films/1/",
			"http://swapi.dev/api/films/3/",
			"http://swapi.dev/api/films/4/",
			"http://swapi.dev/api/films/5/",
			"http://swapi.dev/api/films/6/",
		}}, Planet{
		Name:    "Mon Cala",
		Climate: "temperate",
		Terrain: "oceans, reefs, islands",
		Films:   []string{}}}

	planetInSample := CreatePlanetDB(planets, inputInPlanets)
	if planetInSample.Films != 5 {
		t.Fatalf("Error condition")
	}

	planetWithoutSample := CreatePlanetDB(planets, inputWithoutPlanets)
	if planetWithoutSample.Films != 0 {
		t.Fatalf("Error condition")
	}
}

func TestCheckPlanetExists(t *testing.T) {
	input := CreatePlanetInput{
		Name:    "Tatooine",
		Climate: "arid",
		Terrain: "desert",
	}
	planets := []Planet{Planet{
		Name:    "Tatooine",
		Climate: "arid",
		Terrain: "desert",
		Films: []string{
			"http://swapi.dev/api/films/1/",
			"http://swapi.dev/api/films/3/",
			"http://swapi.dev/api/films/4/",
			"http://swapi.dev/api/films/5/",
			"http://swapi.dev/api/films/6/",
		}}, Planet{
		Name:    "Mon Cala",
		Climate: "temperate",
		Terrain: "oceans, reefs, islands",
		Films:   []string{}}}

	planetTrue := CheckPlanetExists(planets[0], input)
	planetFalse := CheckPlanetExists(planets[1], input)

	if planetTrue == false {
		t.Fatalf("Error condition")
	}
	if planetFalse == true {
		t.Fatalf("Error condition")
	}

}
