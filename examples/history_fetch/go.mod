module github.com/5c077m4n/pikud-haoref-api-go/examples/history_fetch

go 1.26.1

replace (
	github.com/5c077m4n/pikud-haoref-api-go/history => ../../history/
	github.com/5c077m4n/pikud-haoref-api-go/poller => ../../poller/
)

require github.com/5c077m4n/pikud-haoref-api-go/history v0.0.0

require (
	github.com/5c077m4n/pikud-haoref-api-go/poller v0.0.0 // indirect
	github.com/goccy/go-json v0.10.6 // indirect
)
