package logz

import "github.com/rs/zerolog"

type Adapter interface {
	Error(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
}

type AdapterNoop struct{}

func (AdapterNoop) Error(msg string, keysAndValues ...interface{}) {}
func (AdapterNoop) Info(msg string, keysAndValues ...interface{})  {}
func (AdapterNoop) Debug(msg string, keysAndValues ...interface{}) {}
func (AdapterNoop) Warn(msg string, keysAndValues ...interface{})  {}

var _ Adapter = AdapterNoop{}

// AdapterKV fit for msg, keyvalue interface, Ex: retryablehttp.
//
//	myLogFormat := log.With().Str("log_source", "mycomponent").Logger()
//	kvLogger := logz.AdapterKV{Log: myLogFormat}
//	kvLogger.Error("this is message", "err", "failed x")
type AdapterKV struct {
	Log        zerolog.Logger
	FrameCount int
	Caller     bool
}

var _ Adapter = AdapterKV{}

func (l AdapterKV) frameUp() zerolog.Logger {
	count := 3
	if l.FrameCount > 0 {
		count = l.FrameCount
	}

	if l.Caller {
		return l.Log.With().CallerWithSkipFrameCount(count).Logger()
	}

	return l.Log.With().Logger()
}

func (l AdapterKV) Error(msg string, keysAndValues ...interface{}) {
	f := l.frameUp()
	f.Error().Fields(keysAndValues).Msg(msg)
}

func (l AdapterKV) Info(msg string, keysAndValues ...interface{}) {
	f := l.frameUp()
	f.Info().Fields(keysAndValues).Msg(msg)
}

func (l AdapterKV) Debug(msg string, keysAndValues ...interface{}) {
	f := l.frameUp()
	f.Debug().Fields(keysAndValues).Msg(msg)
}

func (l AdapterKV) Warn(msg string, keysAndValues ...interface{}) {
	f := l.frameUp()
	f.Warn().Fields(keysAndValues).Msg(msg)
}
