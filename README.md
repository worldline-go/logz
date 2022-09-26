# logz - Zerolog Helper

This library help to easily initialize log formats in projects.

```sh
go get github.com/worldline-go/logz
```

## Usage

InitializeLog auto format, json or pretty print.  
Use `LOG_PRETTY` boolean env value `(1, t, T, TRUE, true, True)` to set it.

```go
logz.InitializeLog(nil)

log.Info().Msg("Log test 1 2 1 2")
```

Caller disabled by default to enable it set with config.

```go
logz.InitializeLog(logz.DefaultLogSettings.SetCaller(true))
```

To change formats, change logz values before the initialize.

```go
logz.TimeFormat       = time.RFC3339Nano
logz.TimePrettyFormat = "2006-01-02 15:04:05"
```

Results of example `go run -trimpath _example/main.go`

In pretty format

```sh
2022-09-26 17:28:26 INF default ctx log
2022-09-26 17:28:26 INF Log test 1 2 1 2
2022-09-26 17:28:26 ERR this is message err="failed x" log_source=mycomponent
2022-09-26 17:28:26 DBG helloo level info but show debug component=test
```

In container

```sh
{"level":"info","time":"2022-09-26T17:28:33.64382204+02:00","message":"default ctx log"}
{"level":"info","time":"2022-09-26T17:28:33.64385865+02:00","message":"Log test 1 2 1 2"}
{"level":"error","log_source":"mycomponent","err":"failed x","time":"2022-09-26T17:28:33.64387864+02:00","message":"this is message"}
{"level":"info","component":"test","time":"2022-09-26T17:28:33.64388865+02:00","level":"debug","message":"helloo level info but show debug"}
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
