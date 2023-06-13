package handlers

import (
	"currency/internal/core"
	"currency/internal/storage"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetLatest(c echo.Context) error {
	currencies := []string{core.UAH, core.EUR, core.JPY}
	url := fmt.Sprintf("https://api.exchangerate.host/latest?base=%s&symbols=%s", core.USD, strings.Join(currencies[:], ","))
	dt := time.Now()
	date := dt.Format("2006-01-02")
	var rateByDate core.CleanLatestRates
	var err error
	rateByDate, err = storage.GetRecordByDate(date)

	if err != nil {
		rateByDate, err = core.ReadRates(url)
		if err != nil {
			log.Fatal(err)
			panic(err.Error())
		}
		storage.AddRecord(rateByDate)
	}

	out, err := json.Marshal(&rateByDate)

	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	return c.String(http.StatusOK, string(out))
}
func GetByDate(c echo.Context) error {
	currencies := []string{core.UAH, core.EUR, core.JPY}
	url := fmt.Sprintf("https://api.exchangerate.host/latest?base=%s&symbols=%s", core.USD, strings.Join(currencies[:], ","))

	date := c.QueryParam("date")
	rateByDate, err := storage.GetRecordByDate(date)

	if err != nil {
		rateByDate, err = core.ReadRates(url)
		if err != nil {
			log.Fatal(err)
			panic(err.Error())
		}
		storage.AddRecord(rateByDate)
	}

	out, err := json.Marshal(&rateByDate)

	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	return c.String(http.StatusOK, string(out))
}
