package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
	"currency/internal/core/rest"
	"net/http"
	"encoding/json"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetByDate(c echo.Context) error {
	date := c.QueryParam("date")

	rates, err := rest.ReadUsdBased("--", date, rest.UAH, rest.EUR, rest.JPY)

	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	out, jsonErr := json.Marshal(&rates)
	
	if jsonErr != nil {
		panic(jsonErr.Error())
	}
	

	return c.String(http.StatusOK, string(out))
}