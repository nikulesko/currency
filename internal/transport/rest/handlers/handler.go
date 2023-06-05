package handlers

import (
	"currency/internal/core/rest"
	"currency/internal/storage"

	"encoding/json"
	"net/http"
	"time"
	
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetLatest(c echo.Context) error {
	dt := time.Now()
	date := dt.Format("2006-01-02")
	var rateByDate rest.CleanLatestRates
	var err error
	rateByDate, err = storage.GetRecordByDate(date)

	if err != nil {
		rateByDate, err = rest.ReadUsdBased(rest.USD, rest.UAH, rest.EUR, rest.JPY)
		if err != nil {
			panic(err.Error())
		}
		storage.AddRecord(rateByDate)
	}

	out, err := json.Marshal(&rateByDate)

	if err != nil {
		panic(err.Error())
	}

	return c.String(http.StatusOK, string(out))
}
func GetByDate(c echo.Context) error {
	date := c.QueryParam("date")
	rateByDate, err := storage.GetRecordByDate(date)

	if err != nil {
		rateByDate, err = rest.ReadUsdBased(rest.USD, rest.UAH, rest.EUR, rest.JPY)
		if err != nil {
			panic(err.Error())
		}
		storage.AddRecord(rateByDate)
	}

	out, err := json.Marshal(&rateByDate)

	if err != nil {
		panic(err.Error())
	}

	return c.String(http.StatusOK, string(out))
}
