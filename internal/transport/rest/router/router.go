package router

import (
	"currency/internal/transport/rest/handlers"

	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", handlers.Index)
	e.GET("/currency/latest", handlers.GetLatest)
	e.GET("/currency/history", handlers.GetByDate)

	return e
}
