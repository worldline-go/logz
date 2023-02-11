# LogZ for echo middleware

Helpful middleware for echo framework.

```sh
go get github.com/worldline-go/logz/logecho
```

## Usage

```go
e.Use(
    middleware.RequestID(),
    middleware.RequestLoggerWithConfig(logecho.RequestLoggerConfig()),
    logecho.ZerologLogger(),
)
```
