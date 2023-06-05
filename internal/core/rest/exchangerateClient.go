package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func ReadUsdBased(date string, currencies ...string) (CleanLatestRates, error) {
	url := fmt.Sprintf("https://api.exchangerate.host/latest?base=%s&symbols=%s", USD, strings.Join(currencies[:], ","))
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return EmptyLatestRates(), err
	}

	var rates LatestRates

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		return EmptyLatestRates(), err
	}

	err = json.Unmarshal(body, &rates)

	if err != nil {
		fmt.Println(err.Error())
		return EmptyLatestRates(), err
	}

	return rates.Clean(), nil
}
