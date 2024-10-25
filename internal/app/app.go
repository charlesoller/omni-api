package app

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/charlesoller/omni-api/internal/database"
	"github.com/charlesoller/omni-api/internal/db"
	"github.com/charlesoller/omni-api/internal/movie"
	"github.com/labstack/echo/v4"
)

func Setup() *echo.Echo {
	e := echo.New()
	api := e.Group("/api")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	go func() {
		log.Println("Starting pprof server on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil)) // Default pprof endpoints are served here
	}()

	// Database
	data := database.SetupDatabase()
	queries := db.New(data)
	store := database.NewStore(data, queries)

	// Repository
	movieRepository := movie.NewMovieRepository(store)

	// Service
	movieService := movie.NewMovieService(movieRepository)

	// Handler
	movieHandler := movie.NewMovieHandler(movieService)

	// Router
	movieRouter := movie.NewMovieRouter(movieHandler)

	// Route Groups
	movieRoutes := api.Group("/movies")

	// Registering Routes
	movieRouter.RegisterRoutes(movieRoutes) // Register movie routes

	return e
}
