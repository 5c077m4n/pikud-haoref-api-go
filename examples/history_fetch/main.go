package main

import (
	"context"
	"log/slog"

	"github.com/5c077m4n/pikud-haoref-api-go/history"
)

func main() {
	alerts, err := history.FetchAlerts(context.TODO())
	if err != nil {
		panic(err)
	}

	slog.Info("Got results", "value", alerts)
}
