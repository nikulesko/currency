package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func ReadUsdBased(token string, date string, currencies ...string) (HistoricalRates, error) {
	url := fmt.Sprintf("https://api.currencyfreaks.com/v2.0/rates/historical?apikey=%s&date=%s&base=%s&symbols=%s", token, "2023-05-23", USD, strings.Join(currencies[:], ","))
	resp, errGet := http.Get(url)
	if errGet != nil {
		fmt.Println(errGet.Error())
		return EmptyHistoricalRates(), errGet
	}

	var rates HistoricalRates

	body, readErr := io.ReadAll(resp.Body)

	if readErr != nil {
		fmt.Println(readErr.Error())
		return EmptyHistoricalRates(), readErr
	}

	err := json.Unmarshal(body, &rates)

	if err != nil {
		fmt.Println(err.Error())
		return EmptyHistoricalRates(), err
	}

	return rates, nil
}
