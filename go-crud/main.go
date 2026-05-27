package main

import (
	"net/http"
	"log"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/joho/godotenv"

	mydb "go-crud/db"
)

func main() {
	godotenv.Load()

	_, err := mydb.Connect()
	if err != nil {
		log.Fatal(err)
	}

	app := echo.New()
	app.Use(middleware.RequestLogger())

	app.GET("/health", func(c *echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	if err := app.Start(":8080"); err != nil {
		app.Logger.Error("failed to start server", "error", err)
	}
}