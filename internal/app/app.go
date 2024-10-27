package app

import (
	"log"
	"net/http"
	"fmt"
	_ "net/http/pprof"

	"github.com/charlesoller/omni-api/internal/database"
	"github.com/charlesoller/omni-api/internal/db"
	"github.com/charlesoller/omni-api/internal/movie"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup() *echo.Echo {
	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			value, _ := c.Get("customValueFromContext").(int)
			fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v\n", v.URI, v.Status, value)
			return nil
		},
	}))

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
	movieRouter.RegisterRoutes(movieRoutes)

	return e
}
