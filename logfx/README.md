# logfx

zerolog for fx dependency injection library.

```sh
go get github.com/worldline-go/logz/logfx
```

## Usage

```go
fx.WithLogger(func() fxevent.Logger {
    logz.InitializeLog(logz.WithServiceInfo("my-service", "v0.1.0"))
    return logfx.Event(log.Logger)
}),
```

Detailed version with config loader

```go
type Config struct {
	LogLevel string
}

func LoadConfig() *Config {
	return &Config{
		LogLevel: "info",
	}
}

func main() {
	fx.New(
		fx.WithLogger(func(cfg *Config) (fxevent.Logger, error) {
			logz.InitializeLog(logz.WithServiceInfo("my-service", "v0.1.0"))
			if err := logz.SetLogLevel(cfg.LogLevel); err != nil {
				return nil, err
			}

			return logfx.Event(log.Logger), nil
		}),
		fx.Provide(
			LoadConfig,
		),
	).Run()
}
```
