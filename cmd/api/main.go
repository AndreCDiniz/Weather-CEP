package main

import (
	"fmt"
	"github.com/AndreCDiniz/Weather-CEP/internal/clients"
	"github.com/AndreCDiniz/Weather-CEP/internal/handlers"
	"github.com/AndreCDiniz/Weather-CEP/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file")
	}

	fmt.Println("====================================")
	fmt.Println("    Weather-CEP API Starting Up     ")
	fmt.Println("====================================")

	gin.SetMode(gin.ReleaseMode)

	viaCEPClient := clients.NewViaCEPClient()
	weatherClient := clients.NewWeatherClient(os.Getenv("WEATHER_API_KEY"))

	weatherService := services.NewWeatherService(viaCEPClient, weatherClient)

	weatherHandler := handlers.NewWeatherHandler(weatherService)

	r := gin.Default()
	r.GET("/weather/:cep", weatherHandler.GetWeather)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	fmt.Printf("\n🚀 Server is ready to handle requests\n")
	fmt.Printf("📡 HTTP Server listening on port: %s\n", port)
	fmt.Printf("🌐 Access the API at: http://localhost:%s/weather/{cep}\n", port)
	fmt.Printf("💡 Example request: http://localhost:%s/weather/01001000\n\n", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Error starting server: %v", err)
	}
}
