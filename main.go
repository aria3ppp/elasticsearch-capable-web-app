package main

import (
	"encoding/json"
	"os"

	"elasticsearch-capable-web-app/db"
	"elasticsearch-capable-web-app/handler"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	dsn := os.Getenv("APP_DSN")
	logger.Info().Interface("dsn", dsn)
	dbInstance, err := db.New(dsn, logger)
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
	resp, err := esClient.Indices.Create("posts")
	if err != nil {
		logger.Err(err).Msg("Elasticsearch failed creating index")
		os.Exit(1)
	}
	if resp.IsError() {
		var e map[string]any
		if err = json.NewDecoder(resp.Body).Decode(&e); err != nil {
			logger.Err(err).
				Msg("failed decoding elasticsearch response body error")
		} else {
			logger.Error().Interface("Elasticsearch response error", e)
		}
		os.Exit(1)
	}

	h := handler.New(dbInstance, esClient, logger)
	router := echo.New()
	rg := router.Group("/v1")
	h.Register(rg)
	router.Start(":8080")
}
