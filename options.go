package logz

type options struct {
	pretty    *bool
	timeStamp *bool
	caller    *bool
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
