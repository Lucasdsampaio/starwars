package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucasdsampaio/starwars/api"
	"github.com/lucasdsampaio/starwars/models"
)

func main() {
	setupServer().Run()
}

func setupServer() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := gin.Default()
	models.ConnectDataBase()

	app.POST("/planet", api.CreatePlanet)
	app.GET("/planets", api.ListPlanets)
	app.GET("/planet-name/:name", api.GetPlanetByName)
	app.GET("/planet-id/:id", api.GetPlanetByID)
	app.DELETE("/planet/:id", api.DeletePlanet)
	return app
}
