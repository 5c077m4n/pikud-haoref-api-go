# An Unofficial SDK for Pikud HaOref's API

This is an unofficial Go library for interacting with Pikud HaOref (Home Front
Command) API. It provides access to real-time alerts, historical alerts, and
Israeli city data.

> [!NOTE]
> This code will work **only** when running it from inside Israel (or via proxy)

## Installation

```bash
go get github.com/5c077m4n/pikud-haoref-api-go
```

## Packages

### History

The `history` package fetches historical alerts from Pikud HaOref. It provides
two modes of operation:

- **Fetch**: Retrieve all historical alerts at once
- **Stream**: Polls for updates continuously

```go
import "github.com/5c077m4n/pikud-haoref-api-go/history"
```

**Functions:**

- `FetchAlerts()` - Fetches all historical alerts
- `Stream(ctx context.Context)` - Returns a channel that streams alerts

### Cities

The `cities` package provides access to Israeli city data, including city names,
area information, and bunker access times.

```go
import "github.com/5c077m4n/pikud-haoref-api-go/cities"
```

**Functions:**

- `FetchCities()` - Returns all cities with metadata (name, area, bunker time)

### Alerts (Deprecated)

The `alerts` package is deprecated. Use `history` instead.

```go
import "github.com/5c077m4n/pikud-haoref-api-go/alerts"
```

## Examples

- **[Cities](./examples/cities/)** - Fetch all of Israel's city names (using
  cache)
- **[History Fetch](./examples/history_fetch/)** - Fetch historical alerts once
- **[History Stream](./examples/history_stream/)** - Fetch a stream of
  historical alerts
