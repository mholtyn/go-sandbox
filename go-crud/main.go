package main

import (
	"net/http"
	"log"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/joho/godotenv"

	mydb "go-crud/db"
	"go-crud/handler"
	"go-crud/store"
)

func main() {
	godotenv.Load()

	log.Println("connecting...")
	db, err := mydb.Connect()
	if err != nil {
		log.Fatal(err)
	}

	taskStore := store.NewTaskStore(db)
	taskHandler := handler.NewTaskHandler(taskStore)

	app := echo.New()
	app.Use(middleware.RequestLogger())

	app.GET("/health", func(c *echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	taskHandler.RegisterRoutes(app)

	if err := app.Start(":8080"); err != nil {
		app.Logger.Error("failed to start server", "error", err)
	}
}