package movie

import (
	"github.com/labstack/echo/v4"
)

type MovieRouter struct {
	h *MovieHandler
}

func NewMovieRouter(handler *MovieHandler) *MovieRouter {
	return &MovieRouter{
		h: handler,
	}
}

func (r *MovieRouter)RegisterRoutes(e *echo.Group) {
	e.GET("", r.h.GetAllMoviesHandler)
	e.GET("/:id", r.h.GetMovieHandler)
}
