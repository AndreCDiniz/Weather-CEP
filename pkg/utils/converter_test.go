package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidCEP(t *testing.T) {
	tests := []struct {
		name     string
		cep      string
		expected bool
	}{
		{
			name:     "Valid CEP",
			cep:      "01001000",
			expected: true,
		},
		{
			name:     "Invalid CEP - Letters",
			cep:      "0100100a",
			expected: false,
		},
		{
			name:     "Invalid CEP - Short",
			cep:      "0100100",
			expected: false,
		},
		{
			name:     "Invalid CEP - Long",
			cep:      "010010000",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidCep(tt.cep)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTemperatureConversions(t *testing.T) {
	tests := []struct {
		name           string
		celsius        float64
		expectedFahr   float64
		expectedKelvin float64
	}{
		{
			name:           "Zero Celsius",
			celsius:        0,
			expectedFahr:   32,
			expectedKelvin: 273.15,
		},
		{
			name:           "Positive Temperature",
			celsius:        25,
			expectedFahr:   77,
			expectedKelvin: 298.15,
		},
		{
			name:           "Negative Temperature",
			celsius:        -10,
			expectedFahr:   14,
			expectedKelvin: 263.15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fahr := CelsiusToFahrenheit(tt.celsius)
			kelvin := CelsiusToKelvin(tt.celsius)

			assert.InDelta(t, tt.expectedFahr, fahr, 0.1)
			assert.InDelta(t, tt.expectedKelvin, kelvin, 0.1)
		})
	}
}
