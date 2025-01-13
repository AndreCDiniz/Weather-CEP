# Weather CEP API
An API to get temperature information from a Brazilian postal code (CEP).

## Technologies Used
- Go 1.21
- Gin Framework
- Docker
- Google Cloud Run

## Prerequisites
- Docker
- Docker Compose
- Go 1.21 or higher (for local development)
- WeatherAPI account (https://www.weatherapi.com/)

## Setup
1. Clone the repository:
```bash
git clone https://github.com/AndreCDiniz/Weather-CEP.git
cd Weather-CEP
```

2. Set up environment variables:
```bash
cp .env.example .env
```
Edit the `.env` file and add your WeatherAPI key.

## Running Locally
1. With Docker:
```bash
docker-compose up --build
```

2. Without Docker:
```bash
go mod download
go run cmd/api/main.go
```

## Tests
Run the tests:
```bash
go test ./... -v
```

## Endpoints
### GET /weather/:cep
Returns the current temperature for the location of the provided postal code.

**Parameters:**
- cep: Valid 8-digit Brazilian postal code

**Live API URL:**
```
https://weather-app-374566725681.us-central1.run.app/weather/{cep}
```

**Example Request:**
```
https://weather-app-374566725681.us-central1.run.app/weather/01001000
```

**Responses:**
- 200: Success
```json
{
    "temp_C": 28.5,
    "temp_F": 83.3,
    "temp_K": 301.65
}
```
- 422: Invalid postal code
- 404: Postal code not found
- 500: Internal error

## Deployment
The application is deployed and available at:
```
https://weather-app-374566725681.us-central1.run.app
```

## Project Structure
```
weather-cep/
├── cmd/                 # Application entrypoints
│   └── api/
│       └── main.go     # Main application file
├── internal/           # Private application code
│   ├── domain/         # Domain models and interfaces
│   │   └── models/     
│   ├── handlers/       # HTTP request handlers
│   ├── services/       # Business logic
│   └── clients/        # External service clients
└── pkg/               # Public libraries
    └── utils/         # Utility functions
```

## Architecture
The project follows a clean architecture approach with the following layers:
1. **Handlers**: HTTP request handling and response formatting
2. **Services**: Business logic and orchestration
3. **Clients**: External API communication
4. **Models**: Data structures and domain objects

## Development
### Adding New Features
1. Create necessary models in `internal/domain/models`
2. Implement business logic in `internal/services`
3. Add HTTP handlers in `internal/handlers`
4. Update tests accordingly

### Running Tests
```bash
# Run all tests
go test ./... -v
# Run tests with coverage
go test ./... -cover -coverprofile=coverage.out
# View coverage in browser
go tool cover -html=coverage.out
```

## API Documentation
### Success Response Example
```json
{
    "temp_C": 25.0,
    "temp_F": 77.0,
    "temp_K": 298.15
}
```

### Error Response Examples
Invalid postal code:
```json
{
    "message": "invalid zipcode"
}
```

Postal code not found:
```json
{
    "message": "can not find zipcode"
}
```

## Deployment Steps
1. Install Google Cloud CLI
2. Configure the project:
```bash
gcloud init
gcloud auth configure-docker
```

3. Build and push the image:
```bash
gcloud builds submit --tag gcr.io/sunlit-unison-447723-u8/weather-app
```

4. Deploy to Cloud Run:
```bash
gcloud run deploy weather-app --image gcr.io/sunlit-unison-447723-u8/weather-app --platform managed --region us-central1 --allow-unauthenticated --project sunlit-unison-447723-u8
```

## Contributing
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
