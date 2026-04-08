# Fetching historical alerts once

This example demonstrates how to fetch historical alerts a single time using the
`history` package.

## What it does

1. Calls `history.FetchAlerts(context.TODO())` to retrieve all historical alerts
   from the API
2. Returns a slice of `Alert` structs containing:
   - Date: When the alert occurred
   - Title: Alert title
   - City: The affected city
   - Category: Alert type (missiles, earthquake, etc.)
3. Logs the results using structured logging

## Run the example

```bash
cd examples/history_fetch
go run main.go
```
