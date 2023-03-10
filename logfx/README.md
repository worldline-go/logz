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
			logz.InitializeLog(logz.WithCaller(false))
			if err := logz.SetLogLevel(cfg.LogLevel); err != nil {
				return nil, err
			}

			return logfx.Event(log.Logger, logfx.WithServiceMessage("my-service", "v0.1.0")), nil
		}),
		fx.Provide(
			LoadConfig,
		),
	).Run()
}
```
