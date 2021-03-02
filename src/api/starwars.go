package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasdsampaio/starwars/models"
)

func CreatePlanet(c *gin.Context) {
	var input CreatePlanetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseAPI, err := AsyncGet()
	if err != nil {
		c.JSON(http.StatusGatewayTimeout, gin.H{"data": "Bad Gateway"})
		return
	}

	planets, err := GetPlanetsFromResponse(responseAPI)
	if err != nil {
		c.JSON(http.StatusGatewayTimeout, gin.H{"data": "Bad Gateway"})
		return
	}

	planetDB := CreatePlanetDB(planets, input)
	models.DB.Create(&planetDB)
	c.JSON(http.StatusOK, gin.H{"data": planetDB})
}

func ListPlanets(c *gin.Context) {
	var planets []models.Planet
	models.DB.Find(&planets)

	c.JSON(http.StatusOK, gin.H{"data": planets})
}

func GetPlanetByName(c *gin.Context) {
	var planet []models.Planet

	if err := models.DB.Where("name = ?", c.Param("name")).Find(&planet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": planet})
}

func GetPlanetByID(c *gin.Context) {
	var planet models.Planet

	if err := models.DB.Where("id = ?", c.Param("id")).First(&planet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": planet})
}

func DeletePlanet(c *gin.Context) {
	var planet models.Planet

	if err := models.DB.Where("id = ?", c.Param("id")).First(&planet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record not found!"})
		return
	}

	models.DB.Delete(planet)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
