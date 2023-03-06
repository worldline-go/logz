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

To get back the logger from the context:

```go
// this will show the request ID in the log
log.Ctx(c.Request().Context()).Info().Msg("Hello world")
```
