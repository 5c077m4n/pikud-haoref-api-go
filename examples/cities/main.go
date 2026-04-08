package main

import (
	"context"
	"log/slog"
	"sync"

	"github.com/5c077m4n/pikud-haoref-api-go/pikuf-haoref-api-go/cities"
)

var fetchAllCityNamesOnce = sync.OnceValues(func() ([]string, error) {
	allCities, err := cities.FetchCities(context.TODO())
	if err != nil {
		return nil, err
	}

	cityNames := make([]string, 0, len(allCities))
	for _, c := range allCities {
		cityNames = append(cityNames, c.Label)
	}

	return cityNames, nil
})

func main() {
	cityNames, err := fetchAllCityNamesOnce()
	if err != nil {
		panic(err)
	}

	slog.Info("Got all Israel's city names", "value", cityNames)
}
