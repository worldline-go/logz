# logz - Zerolog Helper

This library help to easily initialize log formats in projects.

```sh
go get github.com/worldline-go/logz
```

Check helper for other modules:

__-__ [logfx](./logfx/README.md) -> zerolog for uber fx dependecy injection library  
__-__ [logecho](./logecho/README.md) -> zerolog for echo web framework  

## Usage

InitializeLog auto format, json or pretty print.  
Use `LOG_PRETTY` boolean env value `(1, t, T, TRUE, true, True)` to set it.

```go
logz.InitializeLog()

log.Info().Msg("Log test 1 2 1 2")
```

Caller disabled by default to enable it set with config.

```go
logz.InitializeLog(logz.WithCaller(true))
```

To modify the default logger with adding some additional values:

```go
logz.InitializeLog(
    logz.WithLogContextFunc(func(ctx zerolog.Context) zerolog.Context {
        return ctx.Str("log_source", "main")
    }),
    logz.WithServiceInfo("awesome-service", "v0.2.4"),
)
```

InitializeLog also adds the generated logger to the `DefaultContextLogger`.  
If not found any logger in context, it will return `log.Logger`.

```go
ctx = log.Ctx(ctx).With().Str("component", "context-test").Logger().WithContext(ctx)
log.Ctx(ctx).Info().Msg("testing")
// 2023-03-10 14:37:24 CET INF _example/main.go:38 > testing component=context-test
```

To change formats, change logz values before the initialize.

```go
logz.TimeFormat       = time.RFC3339Nano
logz.TimePrettyFormat = "2006-01-02 15:04:05 MST"
```

Results of example `go run --trimpath _example/main.go`

In pretty format

```sh
2023-03-15 10:11:58 CET INF ./main.go:25 > default ctx log log_source=main service_name=awesome-service service_version=v0.2.4
2023-03-15 10:11:58 CET INF ./main.go:27 > Log test 1 2 1 2 log_source=main service_name=awesome-service service_version=v0.2.4
2023-03-15 10:11:58 CET ERR github.com/worldline-go/logz/adapters.go:49 > this is message err="failed x" log_source=mycomponent
2023-03-15 10:11:58 CET DBG ./main.go:37 > helloo level info but show debug component=test log_source=main service_name=awesome-service service_version=v0.2.4
2023-03-15 10:11:58 CET INF ./main.go:43 > testing component=context-test log_source=main service_name=awesome-service service_version=v0.2.4
```

In container

```sh
{"level":"info","time":"2022-11-24T14:56:00.611277862+01:00","caller":"./main.go:20","message":"default ctx log"}
{"level":"info","time":"2022-11-24T14:56:00.611330401+01:00","caller":"./main.go:22","message":"Log test 1 2 1 2"}
{"level":"error","log_source":"mycomponent","err":"failed x","time":"2022-11-24T14:56:00.611339445+01:00","caller":"github.com/worldline-go/logz/adapters.go:31","message":"this is message"}
{"level":"info","component":"test","time":"2022-11-24T14:56:00.611348632+01:00","caller":"./main.go:32","level":"debug","message":"helloo level info but show debug"}
```

### With LogLevel

Hooks usable for forcing to print level information.  
But still we need to use a level writer to effect enable/disable in log level modes.

```go
logPull := log.Ctx(ctx).With().Str("component", "test").Logger().Hook(logz.Hooks.DebugHook)
_, _ = io.Copy(logz.LevelWriter(&logPull, zerolog.DebugLevel), strings.NewReader("message X"))
```

Example for `echo` webframework with `lecho`

```go
e.Logger = lecho.New(log.Logger)
```
