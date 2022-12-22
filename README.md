# logz - Zerolog Helper

This library help to easily initialize log formats in projects.

```sh
go get github.com/worldline-go/logz
```

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

To change formats, change logz values before the initialize.

```go
logz.TimeFormat       = time.RFC3339Nano
logz.TimePrettyFormat = "2006-01-02 15:04:05 MST"
```

Results of example `go run -trimpath _example/main.go`

In pretty format

```sh
2022-11-24 14:55:00 CET INF _example/main.go:20 > default ctx log
2022-11-24 14:55:00 CET INF _example/main.go:22 > Log test 1 2 1 2
2022-11-24 14:55:00 CET ERR github.com/worldline-go/logz/adapters.go:31 > this is message err="failed x" log_source=mycomponent
2022-11-24 14:55:00 CET DBG _example/main.go:32 > helloo level info but show debug component=test
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

Added info level in message and limit it with info level.

```go
e.Logger = lecho.New(loghelper.LevelWriter(log.Logger.Hook(loghelper.Hooks.InfoHook), zerolog.InfoLevel))
```
