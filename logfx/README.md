# logfx

zerolog for fx dependency injection library.

```sh
go get github.com/worldline-go/logz/logfx
```

## Usage

Simple usage with default logger

```go
fx.WithLogger(logfx.New)
```

With custom logger

```go
fx.WithLogger(func() fxevent.Logger {
    return logfx.New(logfx.WithLogger(log.Logger.With().Str("component", "test").Logger()))
})
```

Detailed version with config loader.

```go
type Config struct {
	LogLevel string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		LogLevel: "info",
	}

    // set loglevel again
	logz.SetLogLevel(cfg.LogLevel)

	return cfg, nil
}

func SetLogger() fxevent.Logger {
	logz.InitializeLog(logz.WithServiceInfo(ServiceName, ServiceVersion))
	logz.SetLogLevel(DefaultLogLevel)

	return logfx.New()
}

func main() {
	fx.New(
		fx.WithLogger(SetLogger),
		fx.Provide(
			LoadConfig,
			// ...
		),
		// ...
	).Run()
}
```
