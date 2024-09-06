package logz

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	EnvPretty        = "LOG_PRETTY"
	EnvLevel         = "LOG_LEVEL"
	TimeFormat       = time.RFC3339Nano
	TimePrettyFormat = "2006-01-02 15:04:05 MST"
)

var LogWriter = zerolog.ConsoleWriter{
	Out: os.Stderr,
	FormatTimestamp: func(i interface{}) string {
		if i == nil {
			return ""
		}

		parse, _ := time.Parse(TimeFormat, i.(string))

		return parse.Format(TimePrettyFormat)
	},
}

// InitializeLog choice between json format or common format.
// LOG_PRETTY boolean environment value always override the decision.
// Override with some option argument.
func InitializeLog(opts ...Option) {
	logger := Logger(opts...)
	log.Logger = logger

	zerolog.DefaultContextLogger = &log.Logger
	zerolog.TimeFieldFormat = TimeFormat
}

func Logger(opts ...Option) zerolog.Logger {
	var o option
	for _, opt := range opts {
		opt(&o)
	}

	var logX zerolog.Context

	if checkPretty(o.Pretty, Default.Pretty) {
		logX = zerolog.New(LogWriter).With()
	} else {
		logX = zerolog.New(os.Stderr).With()
	}

	if checkDefault(o.TimeStamp, Default.TimeStamp) {
		logX = logX.Timestamp()
	}

	if checkDefault(o.Caller, Default.Caller) {
		logX = logX.Caller()
	}

	for _, fn := range o.LogContextFuncs {
		logX = fn(logX)
	}

	logger := logX.Logger()

	// set log level global
	if err := SetLogLevel(checkLevel(o.Level, Default.Level)); err != nil {
		logger.Warn().Err(err).Msg("failed to set log level in initialize")
	}

	return logger
}

// SetLogLevel globally changes zerolog's level.
func SetLogLevel(logLevel string) error {
	zerologLevel, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		return fmt.Errorf("LOG_LEVEL %s, err: %w", logLevel, err)
	}

	zerolog.SetGlobalLevel(zerologLevel)

	return nil
}
