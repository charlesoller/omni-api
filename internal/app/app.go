package app

import (
	_ "net/http/pprof"
	"log"
	"net/http"

	// "github.com/charlesoller/omni-api/internal/database"
	// "github.com/charlesoller/omni-api/internal/db"
	"github.com/labstack/echo/v4"
)

func Setup() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})


	// Database
	// data := database.SetupDatabase()
	// queries := db.New(data)
	// store := database.NewStore(data, queries)


	go func() {
			log.Println("Starting pprof server on :6060")
			log.Println(http.ListenAndServe("localhost:6060", nil)) // Default pprof endpoints are served here
	}()

	// movieImportService.(20)	// Start Index

	return e
}
