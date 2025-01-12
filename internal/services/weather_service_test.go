package services

import (
	"errors"
	"github.com/AndreCDiniz/Weather-CEP/internal/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockViaCEPClient implementa a interface ViaCEPClientInterface
type MockViaCEPClient struct {
	mock.Mock
}

// GetLocation implementa o método da interface
func (m *MockViaCEPClient) GetLocation(cep string) (*models.ViaCEPResponse, error) {
	args := m.Called(cep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.ViaCEPResponse), args.Error(1)
}

// MockWeatherClient implementa a interface WeatherClientInterface
type MockWeatherClient struct {
	mock.Mock
}

// GetTemperature implementa o método da interface
func (m *MockWeatherClient) GetTemperature(city string) (*models.WeatherResponse, error) {
	args := m.Called(city)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.WeatherResponse), args.Error(1)
}

func TestGetTemperatureByCEP(t *testing.T) {
	tests := []struct {
		name          string
		cep           string
		setupMocks    func(*MockViaCEPClient, *MockWeatherClient)
		expectedTemp  *models.TemperatureResponse
		expectedError error
	}{
		{
			name: "Success",
			cep:  "01001000",
			setupMocks: func(viaCEP *MockViaCEPClient, weather *MockWeatherClient) {
				viaCEP.On("GetLocation", "01001000").Return(&models.ViaCEPResponse{
					Localidade: "São Paulo",
					Erro:       false,
				}, nil)

				weather.On("GetTemperature", "São Paulo").Return(&models.WeatherResponse{
					Current: struct {
						TempC float64 `json:"temp_c"`
					}{
						TempC: 25.0,
					},
				}, nil)
			},
			expectedTemp: &models.TemperatureResponse{
				TempC: 25.0,
				TempF: 77.0,
				TempK: 298.15,
			},
			expectedError: nil,
		},
		{
			name:          "Invalid CEP",
			cep:           "invalid",
			setupMocks:    func(viaCEP *MockViaCEPClient, weather *MockWeatherClient) {},
			expectedTemp:  nil,
			expectedError: errors.New("invalid zipcode"),
		},
		{
			name: "CEP Not Found",
			cep:  "00000000",
			setupMocks: func(viaCEP *MockViaCEPClient, weather *MockWeatherClient) {
				viaCEP.On("GetLocation", "00000000").Return(&models.ViaCEPResponse{
					Erro: true,
				}, nil)
			},
			expectedTemp:  nil,
			expectedError: errors.New("can not find zipcode"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockViaCEP := new(MockViaCEPClient)
			mockWeather := new(MockWeatherClient)

			tt.setupMocks(mockViaCEP, mockWeather)

			service := NewWeatherService(mockViaCEP, mockWeather)
			temp, err := service.GetTemperatureByCEP(tt.cep)

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
				assert.Nil(t, temp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, temp)
				assert.InDelta(t, tt.expectedTemp.TempC, temp.TempC, 0.1)
				assert.InDelta(t, tt.expectedTemp.TempF, temp.TempF, 0.1)
				assert.InDelta(t, tt.expectedTemp.TempK, temp.TempK, 0.1)
			}

			mockViaCEP.AssertExpectations(t)
			mockWeather.AssertExpectations(t)
		})
	}
}
