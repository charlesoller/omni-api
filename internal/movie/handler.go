package movie

import (
	"context"
	"net/http"

	"strconv"

	"github.com/charlesoller/omni-api/internal/db"
	"github.com/labstack/echo/v4"
	"github.com/pgvector/pgvector-go"
)

type MovieHandler struct {
	s *MovieService
}

func NewMovieHandler(service *MovieService) *MovieHandler {
	return &MovieHandler{
		s: service,
	}
}

func (h *MovieHandler) GetAllMoviesHandler(c echo.Context) error {
	ctx := context.Background()
	var movies []db.Movie

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	title := c.QueryParam("title")
	if title != "" {
		movies, err = h.s.GetAllMoviesMatchingSearch(ctx, int32(page), int32(limit), title)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	} else {
		movies, err = h.s.GetAllMovies(ctx, int32(page), int32(limit))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}

	return c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) GetMovieHandler(c echo.Context) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	movie, err := h.s.GetMovie(ctx, int32(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) GetRandomMovieHandler(c echo.Context) error {
	ctx := context.Background()

	movie, err := h.s.GetRandomMovie(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusAccepted, movie)
}

func (h *MovieHandler) GetMovieDetailsHandler(c echo.Context) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	movie, err := h.s.GetMovieDetails(ctx, int32(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	c.Response().Header().Set("Pragma", "no-cache")
	c.Response().Header().Set("Expires", "0")

	return c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) GetSimilarMoviesHandler(c echo.Context) error {
	ctx := context.Background()

	v := new(pgvector.Vector)
	err := c.Bind(&v)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	movies, err := h.s.GetSimilarMovies(ctx, v)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, movies)
}
