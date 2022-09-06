package main

import (
	"os"

	"elasticsearch-capable-web-app/db"
	"elasticsearch-capable-web-app/handler"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog"
)

// func runThenExit() {
// 	sqlBoilerDemo()
// 	os.Exit(0)
// }

func main() {
	// runThenExit()

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	dbConfig := db.ConnConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DbName:   os.Getenv("POSTGRES_DB"),
	}
	logger.Info().Interface("config", &dbConfig).Msg("config:")
	dbInstance, err := db.New(dbConfig, logger)
	if err != nil {
		logger.Err(err).Msg("Connection failed")
		os.Exit(1)
	}
	logger.Info().Msg("Database connection established")

	// Connect using default address
	esClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		logger.Err(err).Msg("Connection failed")
		os.Exit(1)
	}

	h := handler.New(dbInstance, esClient, logger)
	router := echo.New()
	rg := router.Group("/v1")
	h.Register(rg)
	router.Start(":8080")
}
