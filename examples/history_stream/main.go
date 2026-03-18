package main

import (
	"context"
	"log/slog"

	"github.com/5c077m4n/pikud-haoref-api-go/history"
)

func main() {
	ctx := context.TODO()
	historyResultsCh, historyErrorsCh := history.Stream(ctx)

	for {
		select {
		case <-ctx.Done():
			return
		case err := <-historyErrorsCh:
			slog.Error("Failed to fetch history", "error", err)
		case alerts := <-historyResultsCh:
			slog.Debug("Received historical alerts", "count", len(alerts))
		}
	}
}
