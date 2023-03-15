package logfx

import (
	"strings"

	"github.com/rs/zerolog"
	"go.uber.org/fx/fxevent"
)

// ZeroLogger is an Fx event logger that logs events to Zerolog.
type ZeroLogger struct {
	Logger zerolog.Logger

	options options
}

func Event(logger zerolog.Logger, opts ...Option) fxevent.Logger {
	o := options{
		StartMessage: "started",
		StopMessage:  "stopped",
	}
	for _, opt := range opts {
		opt(&o)
	}

	return &ZeroLogger{
		Logger:  logger,
		options: o,
	}
}

// LogEvent logs the given event to the provided Zap logger.
func (l *ZeroLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Logger.Debug().Str("callee", e.FunctionName).Str("caller", e.CallerName).Msg("OnStart hook executing")
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Logger.Debug().Str("callee", e.FunctionName).Str("caller", e.CallerName).Err(e.Err).Msg("OnStart hook failed")
		} else {
			l.Logger.Debug().Str("callee", e.FunctionName).Str("caller", e.CallerName).Str("runtime", e.Runtime.String()).Msg("OnStart hook executed")
		}
	case *fxevent.OnStopExecuting:
		l.Logger.Debug().Str("callee", e.FunctionName).Str("caller", e.CallerName).Msg("OnStop hook executing")
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Logger.Debug().Str("callee", e.FunctionName).Str("caller", e.CallerName).Err(e.Err).Msg("OnStop hook failed")
		} else {
			l.Logger.Debug().Str("callee", e.FunctionName).Str("caller", e.CallerName).Str("runtime", e.Runtime.String()).Msg("OnStop hook executed")
		}
	case *fxevent.Supplied:
		if e.Err != nil {
			moduleField(e.ModuleName, l.Logger.Debug().Str("type", e.TypeName)).Err(e.Err).Msg("error encountered while applying options")
		} else {
			moduleField(e.ModuleName, l.Logger.Debug().Str("type", e.TypeName)).Msg("supplied")
		}
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			maybeBool("private", e.Private, moduleField(e.ModuleName, l.Logger.Debug().Str("constructor", e.ConstructorName)).Str("type", rtype)).Msg("provided")
		}
		if e.Err != nil {
			moduleField(e.ModuleName, l.Logger.Debug().Str("constructor", e.ConstructorName)).Err(e.Err).Msg("error encountered while applying options")
		}
	case *fxevent.Replaced:
		for _, rtype := range e.OutputTypeNames {
			moduleField(e.ModuleName, l.Logger.Debug()).Str("type", rtype).Msg("replaced")
		}
		if e.Err != nil {
			moduleField(e.ModuleName, l.Logger.Debug()).Err(e.Err).Msg("error encountered while replacing")
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			moduleField(e.ModuleName, l.Logger.Debug().Str("decorator", e.DecoratorName)).Str("type", rtype).Msg("decorated")
		}
		if e.Err != nil {
			moduleField(e.ModuleName, l.Logger.Debug().Str("decorator", e.DecoratorName)).Err(e.Err).Msg("error encountered while applying options")
		}
	case *fxevent.Invoking:
		// Do not log stack as it will make logs hard to read.
		moduleField(e.ModuleName, l.Logger.Debug().Str("function", e.FunctionName)).Msg("invoking")
	case *fxevent.Invoked:
		if e.Err != nil {
			moduleField(e.ModuleName, l.Logger.Debug().Str("function", e.FunctionName)).Err(e.Err).Msg("invocation failed")
		}
	case *fxevent.Stopping:
		l.Logger.Warn().Str("signal", strings.ToUpper(e.Signal.String())).Msg("received signal")
	case *fxevent.Stopped:
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).Msgf("stop failed %s", l.options.AppendMessage)
		} else {
			l.Logger.Info().Msgf("%s %s", l.options.StopMessage, l.options.AppendMessage)
		}
	case *fxevent.RollingBack:
		l.Logger.Info().Err(e.StartErr).Msg("start failed, rolling back")
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).Msg("rollback failed")
		}
	case *fxevent.Started:
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).Msgf("start failed %s", l.options.AppendMessage)
		} else {
			l.Logger.Info().Msgf("%s %s", l.options.StartMessage, l.options.AppendMessage)
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).Msg("custom logger initialization failed")
		} else {
			l.Logger.Debug().Str("function", e.ConstructorName).Msg("initialized custom fxevent.Logger")
		}
	}
}

func moduleField(name string, z *zerolog.Event) *zerolog.Event {
	if len(name) == 0 {
		return z
	}
	return z.Str("module", name)
}

func maybeBool(name string, b bool, z *zerolog.Event) *zerolog.Event {
	if b {
		return z.Bool(name, b)
	}

	return z
}
