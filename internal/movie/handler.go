package movie

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"strconv"
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

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1 
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	movies, err := h.s.GetAllMovies(ctx, int32(page), int32(limit))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, movies)
}
