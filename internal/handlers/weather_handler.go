package handlers

import (
	"github.com/AndreCDiniz/Weather-CEP/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WeatherHandler struct {
	service *services.WeatherService
}

func NewWeatherHandler(service *services.WeatherService) *WeatherHandler {
	return &WeatherHandler{service}
}

func (h *WeatherHandler) GetWeather(c *gin.Context) {
	cep := c.Param("cep")

	temp, err := h.service.GetTemperatureByCEP(cep)
	if err != nil {
		switch err.Error() {
		case "invalid zipcode":
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		case "zipcode not find zipcode":
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, temp)
}
