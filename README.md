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

To change formats, change logz values before the initialize.

```go
logz.TimeFormat       = time.RFC3339Nano
logz.TimePrettyFormat = "2006-01-02 15:04:05"
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
