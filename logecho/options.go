package logecho

import "github.com/rs/zerolog"

type options struct {
	loggerUse bool
	level     zerolog.Level
	logger    zerolog.Logger
}

type Option func(options *options)

// WithLevel sets the log level for the logger, default is debug.
func WithLevel(level zerolog.Level) Option {
	return func(options *options) {
		options.level = level
	}
}

// WithLogger sets the logger to use, default is log.Logger.
func WithLogger(logger zerolog.Logger) Option {
	return func(options *options) {
		options.loggerUse = true
		options.logger = logger
	}
}
