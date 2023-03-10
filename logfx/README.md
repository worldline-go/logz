# logfx

zerolog for fx dependency injection library.

```sh
go get github.com/worldline-go/logz/logfx
```

## Usage

```go
fx.WithLogger(func() fxevent.Logger {
    logz.InitializeLog(logz.WithCaller(false))
    return logfx.Event(log.Logger, logfx.WithServiceMessage("my-service", "v0.1.0"))
}),
```
