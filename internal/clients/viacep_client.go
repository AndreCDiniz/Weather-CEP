package clients

import (
	"encoding/json"
	"fmt"
	"github.com/AndreCDiniz/Weather-CEP/internal/domain/models"
	"net/http"
)

type ViaCEPClientInterface interface {
	GetLocation(cep string) (*models.ViaCEPResponse, error)
}

type ViaCEPClient struct {
	baseURL string
}

func NewViaCEPClient() ViaCEPClientInterface {
	return &ViaCEPClient{
		baseURL: "https://viacep.com.br/ws",
	}
}

func (c *ViaCEPClient) GetLocation(cep string) (*models.ViaCEPResponse, error) {
	url := fmt.Sprintf("%s/%s/json", c.baseURL, cep)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var location models.ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}

	return &location, nil
}
