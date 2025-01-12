package clients

import (
	"encoding/json"
	"fmt"
	"github.com/AndreCDiniz/Weather-CEP/internal/domain/models"
	"net/http"
)

type WeatherClientInterface interface {
	GetTemperature(city string) (*models.WeatherResponse, error)
}

type WeatherClient struct {
	baseURL string
	apiKey  string
}

func NewWeatherClient(apiKey string) WeatherClientInterface {
	return &WeatherClient{
		baseURL: "http://api.weatherapi.com/v1",
		apiKey:  apiKey,
	}
}

func (c *WeatherClient) GetTemperature(city string) (*models.WeatherResponse, error) {
	url := fmt.Sprintf("%s/current.json?key=%s&q=%s", c.baseURL, c.apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weather models.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}

	return &weather, nil
}
