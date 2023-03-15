package logfx

import "github.com/rs/zerolog"

type options struct {
	StartMessage  string
	StopMessage   string
	AppendMessage string
	Logger        zerolog.Logger
}

// Option is a functional option for configuring the logger.
type Option func(*options)

// WithLogger sets the logger to use.
func WithLogger(logger zerolog.Logger) Option {
	return func(o *options) {
		o.Logger = logger
	}
}

// WithStartMessage sets the message to log when the application starts.
func WithStartMessage(msg string) Option {
	return func(o *options) {
		o.StartMessage = msg
	}
}

// WithStopMessage sets the message to log when the application stops.
func WithStopMessage(msg string) Option {
	return func(o *options) {
		o.StopMessage = msg
	}
}

// WithAppendMessage sets the message to log when the application start/stops.
func WithAppendMessage(msg string) Option {
	return func(o *options) {
		o.AppendMessage = " " + msg
	}
}
