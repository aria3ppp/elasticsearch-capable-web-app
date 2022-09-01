package handler

import (
	"elasticsearch-capable-web-app/db"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type Handler struct {
	db       *db.Database
	logger   zerolog.Logger
	esClient *elasticsearch.Client
}

func New(
	database *db.Database,
	esClient *elasticsearch.Client,
	logger zerolog.Logger,
) *Handler {
	return &Handler{
		db:       database,
		esClient: esClient,
		logger:   logger,
	}
}

func (h *Handler) Register(group *echo.Group) {
	group.GET("/posts/:id", h.GetPost)
	group.PATCH("/posts/:id", h.UpdatePost)
	group.DELETE("/posts/:id", h.DeletePost)

	group.GET("/posts", h.GetPosts)
	group.POST("/posts", h.CreatePost)

	group.GET("/search", h.SearchPosts)
}
