package logz

import "github.com/rs/zerolog"

type options struct {
	pretty          *bool
	timeStamp       *bool
	caller          *bool
	level           *string
	logContextFuncs []func(zerolog.Context) zerolog.Context
}

type Option func(options *options)

func WithTimeStamp(timeStamp bool) Option {
	return func(options *options) {
		options.timeStamp = &timeStamp
	}
}

func WithCaller(caller bool) Option {
	return func(options *options) {
		options.caller = &caller
	}
}

func WithPretty(pretty bool) Option {
	return func(options *options) {
		options.pretty = &pretty
	}
}

func WithLevel(level string) Option {
	return func(options *options) {
		options.level = &level
	}
}

func WithLogContextFunc(fn func(zerolog.Context) zerolog.Context) Option {
	return func(options *options) {
		options.logContextFuncs = append(options.logContextFuncs, fn)
	}
}

func WithServiceInfo(serviceName, serviceVersion string) Option {
	return func(options *options) {
		options.logContextFuncs = append(options.logContextFuncs, func(ctx zerolog.Context) zerolog.Context {
			return ctx.Str("service_name", serviceName).Str("service_version", serviceVersion)
		})
	}
}
