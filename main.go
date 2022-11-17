package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	// Connect to elasticsearch
	esClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		logger.Err(err).Msg("Connection failed")
		os.Exit(1)
	}
	// check posts index exists
	postsIndex := os.Getenv("ELASTICSEARCH_INDEX_POSTS")
	resp, err := esClient.Indices.Exists([]string{postsIndex})
	if err != nil {
		logger.Err(err).
			Str("index", postsIndex).
			Msg("esClient.Indices.Exists error")
		return
	}
	if resp.StatusCode == http.StatusNotFound {
		// create posts index
		resp, err := esClient.Indices.Create(postsIndex)
		if err != nil {
			logger.Err(err).
				Str("index", postsIndex).
				Msg("esClient.Indices.Create error")
			return
		}
		if resp.IsError() {
			var e map[string]any
			if err = json.NewDecoder(resp.Body).Decode(&e); err != nil {
				logger.Err(err).
					Str("index", postsIndex).
					Msg("failed decoding elasticsearch response body error")
			} else {
				logger.Error().
					Int("http code", resp.StatusCode).
					Interface("type", e["error"].(map[string]interface{})["type"]).
					Interface("reason", e["error"].(map[string]interface{})["reason"]).
					Msg("esClient.Indices.Create response error")
			}
			return
		}
	} else if resp.IsError() {
		var e map[string]any
		if err = json.NewDecoder(resp.Body).Decode(&e); err != nil {
			logger.Err(err).
				Str("index", postsIndex).
				Msg("failed decoding elasticsearch response body error")
		} else {
			logger.Error().
				Int("http code", resp.StatusCode).
				Interface("type", e["error"].(map[string]interface{})["type"]).
				Interface("reason", e["error"].(map[string]interface{})["reason"]).
				Msg("esClient.Indices.Exists response error")
		}
		return
	}

	h := handler.New(dbInstance, esClient, logger)
	router := echo.New()
	rg := router.Group("/v1")
	h.Register(rg)
	router.Start(":8080")
}
