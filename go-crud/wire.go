//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v5"

	mydb "go-crud/db"
	"go-crud/handler"
	"go-crud/store"
)

func InitializeApp() (*echo.Echo, error) {
	wire.Build(
		mydb.Connect,
		store.NewTaskStore,
		handler.NewTaskHandler,
		NewEcho,
	)
	return nil, nil
}
