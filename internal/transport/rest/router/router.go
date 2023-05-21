package router

import (
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
	"currency/internal/transport/rest/handlers"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", handlers.Index)
	e.GET("/currency", handlers.GetByDate)

	return e
}