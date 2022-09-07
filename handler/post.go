package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"elasticsearch-capable-web-app/db"
	"elasticsearch-capable-web-app/models"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreatePost(c echo.Context) error {
	var post models.Post

	if err := (&echo.DefaultBinder{}).BindBody(c, &post); err != nil {
		h.logger.Err(err).Msg("could not parse request body")
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"error": "invalid request body: " + err.Error(),
		})
	}

	if err := h.db.CreatePost(&post); err != nil {
		h.logger.Err(err).Msg("could not create post")
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"error": "could not save post: " + err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{"post": post})
}

func (h *Handler) UpdatePost(c echo.Context) error {
	var id uint64
	var post models.Post
	var err error
	if id, err = strconv.ParseUint(c.Param("id"), 10, 0); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			echo.Map{"error": "invalid post id"},
		)
	}
	if err = (&echo.DefaultBinder{}).BindBody(c, &post); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			echo.Map{"error": "could not parse request: " + err.Error()},
		)
	}

	err = h.db.UpdatePost(int(id), post)
	if err != nil {
		if err == db.ErrNoRecord {
			return echo.NewHTTPError(
				http.StatusNotFound,
				echo.Map{
					"error": "could not find post with id: " + strconv.FormatUint(
						uint64(id),
						10,
					),
				},
			)
		}
		h.logger.Err(err).Msg("could not update post")
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			echo.Map{
				"error": "could not update post: " + err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, echo.Map{"post": post})
}

func (h *Handler) DeletePost(c echo.Context) error {
	var id uint64
	var err error
	if id, err = strconv.ParseUint(c.Param("id"), 10, 0); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			echo.Map{"error": "invalid post id"},
		)
	}
	err = h.db.DeletePost(int(id))
	if err != nil {
		if err == db.ErrNoRecord {
			return echo.NewHTTPError(
				http.StatusNotFound,
				echo.Map{
					"error": "could not find post with id: " + strconv.FormatUint(
						uint64(id),
						10,
					),
				},
			)
		}

		return echo.NewHTTPError(
			http.StatusInternalServerError,
			echo.Map{"error": err},
		)
	}

	return echo.NewHTTPError(
		http.StatusOK,
		echo.Map{"data": map[string]string{"message": "post deleted"}},
	)
}

func (h *Handler) GetPosts(c echo.Context) error {
	posts, err := h.db.GetPosts()
	if err != nil {
		h.logger.Err(err).Msg("Could not fetch posts")
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}
	return echo.NewHTTPError(http.StatusOK, echo.Map{"data": posts})
}

func (h *Handler) GetPost(c echo.Context) error {
	var id uint64
	var err error
	if id, err = strconv.ParseUint(c.Param("id"), 10, 0); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			echo.Map{"error": "invalid post id"},
		)
	}
	post, err := h.db.GetPostById(int(id))
	if err != nil {
		if err == db.ErrNoRecord {
			return echo.NewHTTPError(
				http.StatusNotFound,
				echo.Map{
					"error": "could not find post with id: " + strconv.FormatUint(
						uint64(id),
						10,
					),
				},
			)
		}
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			echo.Map{"error": err},
		)
	}

	return echo.NewHTTPError(http.StatusOK, echo.Map{"data": post})
}

func (h *Handler) SearchPosts(c echo.Context) error {
	query := c.QueryParam("q")
	if query == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			echo.Map{"error": "no search query presented"},
		)
	}

	body := fmt.Sprintf(
		`{"query": {"multi_match": {"query": "%s", "fields": ["title", "body"], "fuzziness": "AUTO"}}}`,
		query,
	)

	res, err := h.esClient.Search(
		h.esClient.Search.WithContext(context.Background()),
		h.esClient.Search.WithIndex("posts"),
		h.esClient.Search.WithBody(strings.NewReader(body)),
		h.esClient.Search.WithPretty(),
	)
	if err != nil {
		h.logger.Err(err).Msg("elasticsearch error")
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]any
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			h.logger.Err(err).Msg("error parsing the response body")
		} else {
			h.logger.Err(fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)).Msg("failed to search query")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"error": e["error"].(map[string]interface{})["reason"],
		})
	}

	h.logger.Info().Interface("res", res.Status())

	var r map[string]any
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		h.logger.Err(err).Msg("elasticsearch error")
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": r["hits"]})
}
