package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"elasticsearch-capable-web-app/db"
	"elasticsearch-capable-web-app/handler"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Connect to database
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
	// Create posts index if not exists
	postsIndex := os.Getenv("ELASTICSEARCH_INDEX_POSTS")
	postsMappings := `{
		"mappings": {
			"properties": {
				"id": { "type": "keyword", "index": false },
				"title": { "type": "text" },
				"body": { "type": "text" },
				"contributed_by": { "type": "keyword", "index": false },
				"contributed_at": { "type": "date", "index": false },
				"deleted": { "type": "boolean", "index": false }
			}
		}
	}`
	if err := createIndexIfNotExists(esClient, postsIndex, postsMappings); err != nil {
		logger.Err(err).
			Str("index", postsIndex).
			Str("mappings", postsMappings).
			Msg("Creating index failed")
		os.Exit(1)
	}

	// Initialize ans register handlers
	h := handler.New(dbInstance, esClient, logger)
	router := echo.New()
	h.Register(router.Group("/v1"))

	// Start server
	router.Start(":8080")
}

func createIndexIfNotExists(
	client *elasticsearch.Client,
	index string,
	mappings string,
) error {
	// check index exists
	resp, err := client.Indices.Exists([]string{index})
	if err != nil {
		return err
	}
	// create index with mappings
	if resp.StatusCode == http.StatusNotFound {
		resp, err := client.Indices.Create(
			index,
			client.Indices.Create.WithBody(strings.NewReader(mappings)),
		)
		if err != nil {
			return err
		}
		if resp.IsError() {
			return responseError(resp)
		}
	} else if resp.IsError() {
		return responseError(resp)
	}
	return nil
}

func responseError(resp *esapi.Response) error {
	var em map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&em); err != nil {
		return err
	}
	return fmt.Errorf("[%s] %s: %s",
		resp.Status(),
		em["error"].(map[string]interface{})["type"],
		em["error"].(map[string]interface{})["reason"])
}
