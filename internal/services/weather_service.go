package services

import (
	"errors"
	"github.com/AndreCDiniz/Weather-CEP/internal/clients"
	"github.com/AndreCDiniz/Weather-CEP/internal/domain/models"
	"github.com/AndreCDiniz/Weather-CEP/pkg/utils"
)

type WeatherService struct {
	viaCEPClient  clients.ViaCEPClientInterface
	weatherClient clients.WeatherClientInterface
}

func NewWeatherService(viaCEPClient clients.ViaCEPClientInterface, weatherClient clients.WeatherClientInterface) *WeatherService {
	return &WeatherService{
		viaCEPClient:  viaCEPClient,
		weatherClient: weatherClient,
	}
}

func (s *WeatherService) GetTemperatureByCEP(cep string) (*models.TemperatureResponse, error) {
	if !utils.IsValidCep(cep) {
		return nil, errors.New("invalid zipcode")
	}

	location, err := s.viaCEPClient.GetLocation(cep)
	if err != nil {
		return nil, err
	}

	if location.Erro || location.Localidade == "" {
		return nil, errors.New("can not find zipcode")
	}

	weather, err := s.weatherClient.GetTemperature(location.Localidade)
	if err != nil {
		return nil, err
	}

	tempC := weather.Current.TempC
	return &models.TemperatureResponse{
		TempC: tempC,
		TempF: utils.CelsiusToFahrenheit(tempC),
		TempK: utils.CelsiusToKelvin(tempC),
	}, nil
}
