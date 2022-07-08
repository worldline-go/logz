package logz

import "github.com/rs/zerolog"

// AdapterKV fit for msg, keyvalue interface, Ex: retryablehttp.
//
//   myLogFormat := log.With().Str("log_source", "mycomponent").Logger()
//   kvLogger := logz.AdapterKV{Log: myLogFormat}
//   kvLogger.Error("this is message", "err", "failed x")
type AdapterKV struct {
	Log zerolog.Logger
}

func (l AdapterKV) Error(msg string, keysAndValues ...interface{}) {
	l.Log.Error().Fields(keysAndValues).Msg(msg)
}

func (l AdapterKV) Info(msg string, keysAndValues ...interface{}) {
	l.Log.Info().Fields(keysAndValues).Msg(msg)
}

func (l AdapterKV) Debug(msg string, keysAndValues ...interface{}) {
	l.Log.Debug().Fields(keysAndValues).Msg(msg)
}

func (l AdapterKV) Warn(msg string, keysAndValues ...interface{}) {
	l.Log.Warn().Fields(keysAndValues).Msg(msg)
}
