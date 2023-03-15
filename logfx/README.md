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

Detailed version with config loader, we get the config for changing the log level before to fx show the logs.

```go
type Config struct {
	LogLevel string
}

func LoadConfig() (*Config, error) {
	logz.InitializeLog(logz.WithServiceInfo("my-service", "v0.1.0"))

	cfg := &Config{
		LogLevel: "info",
	}

	return cfg, nil
}

func SetLogger(cfg *Config) (fxevent.Logger, error) {
	if err := logz.SetLogLevel(cfg.LogLevel); err != nil {
		return nil, err
	}

	log.Info().Object("config", igconfig.Printer{Value: cfg}).Msg("loaded config")

	return logfx.New(), nil
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
