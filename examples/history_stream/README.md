# Fetching a stream of historical alerts

This example demonstrates how to continuously poll for historical alerts using
the `history` package with streaming.

## What it does

1. Creates a context (can be cancelled to stop streaming)
2. Calls `history.Stream(ctx)` which returns two channels:
   - `historyResultsCh`: Receives slices of `Alert` structs
   - `historyErrorsCh`: Receives errors if the fetch fails
3. Uses a select loop to handle incoming data and errors
4. Polls the API every 2 seconds for updates

## Run the example

```bash
cd examples/history_stream
go run main.go
```

