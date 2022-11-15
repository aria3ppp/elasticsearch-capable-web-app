package main

import (
	"encoding/json"
	"fmt"
	"os"

	"elasticsearch-capable-web-app/db"
	"elasticsearch-capable-web-app/handler"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
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
	resp, err := esClient.Indices.Create(os.Getenv("ELASTICSEARCH_INDEX_POSTS"))
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
