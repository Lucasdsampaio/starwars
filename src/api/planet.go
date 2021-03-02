package api

import (
	"encoding/json"
	"time"

	"github.com/lucasdsampaio/starwars/models"
)

func GetPlanetsFromResponse(responseAPI []*HttpResponse) ([]Planet, error) {
	var planets []Planet
	for _, response := range responseAPI {
		var bodyResponseAPI ResponseAPI

		if err := json.NewDecoder(response.Resp.Body).Decode(&bodyResponseAPI); err != nil {
			panic("Bad Gateway")
		}

		for _, planet := range bodyResponseAPI.Results {
			planets = append(planets, planet)
		}
	}
	return planets, nil
}

func CreatePlanetDB(planets []Planet, input CreatePlanetInput) models.Planet {
	for _, planet := range planets {
		if CheckPlanetExists(planet, input) {
			return models.Planet{Name: planet.Name, Climate: planet.Climate, Terrain: planet.Terrain, Films: len(planet.Films), CreateAt: time.Now()}
		}
	}
	return models.Planet{Name: input.Name, Climate: input.Climate, Terrain: input.Terrain, Films: 0, CreateAt: time.Now()}
}

func CheckPlanetExists(planet Planet, input CreatePlanetInput) bool {
	if planet.Name == input.Name && planet.Climate == input.Climate && planet.Terrain == input.Terrain {
		return true
	}
	return false
}
