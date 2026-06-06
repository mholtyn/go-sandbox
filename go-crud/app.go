package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"go-crud/handler"
)

func NewEcho(h *handler.TaskHandler) *echo.Echo {
	app := echo.New()
	app.Use(middleware.RequestLogger())

	app.GET("/health", func(c *echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	h.RegisterRoutes(app)

	return app
}
