package api

import "net/http"

type CreatePlanetInput struct {
	Name    string `json:"name" binding:"required"`
	Climate string `json:"climate" binding:"required"`
	Terrain string `json:"terrain" binding:"required"`
}

type HttpResponse struct {
	ID   int
	Resp *http.Response
	Err  error
}

type Planet struct {
	Name    string
	Climate string
	Terrain string
	Films   []string
}
type ResponseAPI struct {
	Results []Planet
}
