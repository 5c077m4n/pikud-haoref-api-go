# How to fetch all of Israel's city names (using cache)

This example demonstrates how to fetch all city names in Israel using the
`cities` package with caching.

## What it does

1. Uses `sync.OnceValues` to cache the result of fetching cities
2. Calls `cities.FetchCities(context.TODO())` which returns a list of `CityData`
   structs
3. Extracts city labels and returns them as a slice of strings
4. Logs the result using structured logging

## Run the example

```bash
cd examples/cities
go run main.go
```
