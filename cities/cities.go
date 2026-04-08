// Package cities to fetch all cities in Israel
package cities

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

func FetchCities(ctx context.Context) ([]CityData, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://alerts-history.oref.org.il/Shared/Ajax/GetDistricts.aspx",
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Error(
				"Could not close request body",
				"error", err,
			)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"unexpected status code from Pikud HaOref Districts API %d",
			resp.StatusCode,
		)
	}

	var cities []CityData
	if err := json.NewDecoder(resp.Body).DecodeContext(ctx, &cities); err != nil {
		return nil, err
	}

	return cities, nil
}
