package logz

import (
	"io"

	"github.com/rs/zerolog"
)

type LogHook struct {
	Level string
}

func (h LogHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Str("level", h.Level)
}

// Hooks pass level informations to loggers.
var Hooks = zerolog.LevelHook{
	DebugHook:   LogHook{Level: zerolog.LevelDebugValue},
	InfoHook:    LogHook{Level: zerolog.LevelInfoValue},
	WarnHook:    LogHook{Level: zerolog.LevelWarnValue},
	ErrorHook:   LogHook{Level: zerolog.LevelErrorValue},
	FatalHook:   LogHook{Level: zerolog.LevelFatalValue},
	PanicHook:   LogHook{Level: zerolog.LevelPanicValue},
	TraceHook:   LogHook{Level: zerolog.LevelTraceValue},
	NoLevelHook: LogHook{Level: ""},
}

// LevelWriter function eliminate logger based on level with manually.
// This is usable for writer loggers.
func LevelWriter(e zerolog.Logger, lvl zerolog.Level) io.Writer {
	if lvl < zerolog.GlobalLevel() {
		return io.Discard
	}

	return e
}
