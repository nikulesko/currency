package core

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ReadRates(url string) (CleanLatestRates, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return EmptyLatestRates(), err
	}

	var rates LatestRates

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return EmptyLatestRates(), err
	}

	err = json.Unmarshal(body, &rates)

	if err != nil {
		log.Fatal(err)
		return EmptyLatestRates(), err
	}

	return rates.Clean(), nil
}
