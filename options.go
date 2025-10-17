package logz

import "github.com/rs/zerolog"

type option struct {
	Pretty          *bool
	TimeStamp       *bool
	Caller          *bool
	Level           *string
	NoColor         *bool
	LogContextFuncs []func(zerolog.Context) zerolog.Context
}

type Option func(options *option)

func ReadOptions(opts ...Option) option {
	var option option
	for _, opt := range opts {
		opt(&option)
	}

	return option
}

func WithOption(opt option) Option {
	return func(option *option) {
		*option = opt
	}
}

func WithTimeStamp(timeStamp bool) Option {
	return func(option *option) {
		option.TimeStamp = &timeStamp
	}
}

func WithCaller(caller bool) Option {
	return func(option *option) {
		option.Caller = &caller
	}
}

func WithPretty(pretty bool) Option {
	return func(option *option) {
		option.Pretty = &pretty
	}
}

func WithLevel(level string) Option {
	return func(option *option) {
		option.Level = &level
	}
}

func WithLogContextFunc(fn func(zerolog.Context) zerolog.Context) Option {
	return func(option *option) {
		option.LogContextFuncs = append(option.LogContextFuncs, fn)
	}
}

func WithServiceInfo(serviceName, serviceVersion string) Option {
	return func(option *option) {
		option.LogContextFuncs = append(option.LogContextFuncs, func(ctx zerolog.Context) zerolog.Context {
			return ctx.Str("service_name", serviceName).Str("service_version", serviceVersion)
		})
	}
}

func WithNoColor(noColor bool) Option {
	return func(option *option) {
		option.NoColor = &noColor
	}
}
