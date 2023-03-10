package logfx

type options struct {
	StartMessage  string
	StopMessage   string
	AppendMessage string
}

// Option is a functional option for configuring the logger.
type Option func(*options)

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
		o.AppendMessage = msg
	}
}

// WithServiceMessage sets the message to log when the application start/stops.
//
// This is a convenience function that appends the service name and version to the append message.
func WithServiceMessage(serviceName, serviceVersion string) Option {
	return func(o *options) {
		o.AppendMessage = serviceName + " [" + serviceVersion + "]"
	}
}
