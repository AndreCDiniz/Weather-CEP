# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git and dependencies for testing
RUN apk add --no-cache git gcc musl-dev

# Copy only the go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Run tests before building
RUN go test -v ./...

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /weather-app ./cmd/api

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /weather-app .
COPY .env .env

ENV GIN_MODE=release
ENV PORT=8000

EXPOSE 8000
CMD ["./weather-app"]