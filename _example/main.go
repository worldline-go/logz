package main

import (
	"context"
	"io"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/worldline-go/logz"
)

func main() {
	logz.InitializeLog(
		logz.WithLogContextFunc(func(ctx zerolog.Context) zerolog.Context {
			return ctx.Str("log_source", "main")
		}),
		logz.WithServiceInfo("awesome-service", "v0.2.4"),
	)

	logz.SetLogLevel("info")

	ctx := context.Background()

	log.Ctx(context.Background()).Info().Msg("default ctx log")

	log.Info().Msg("Log test 1 2 1 2")

	myLogFormat := logz.Logger(logz.WithCaller(true)).With().Str("log_source", "mycomponent").Logger()
	kvLogger := logz.AdapterKV{Log: myLogFormat}

	kvLogger.Error("this is message", "err", "failed x")

	// force log level text as debug
	logTest := log.Ctx(ctx).With().Str("component", "test").Logger().Hook(logz.Hooks.DebugHook)
	// use info level for write but it will show debug
	logTest.Info().Msg("helloo level info but show debug")

	// use level with io.Copy
	_, _ = io.Copy(logz.LevelWriter(logTest, zerolog.DebugLevel), strings.NewReader("message X"))

	ctx = log.Ctx(ctx).With().Str("component", "context-test").Logger().WithContext(ctx)
	log.Ctx(ctx).Info().Msg("testing")
}
