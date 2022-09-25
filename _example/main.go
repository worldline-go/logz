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
	logz.InitializeLog(nil)

	logz.SetLogLevel("info")

	ctx := context.Background()

	log.Ctx(context.Background()).Info().Msg("default ctx log")

	log.Info().Msg("Log test 1 2 1 2")

	myLogFormat := logz.GetLogger(logz.DefaultLogSettings.SetCaller(false)).With().Str("log_source", "mycomponent").Logger()
	kvLogger := logz.AdapterKV{Log: myLogFormat, Caller: true}

	kvLogger.Error("this is message", "err", "failed x")

	// force log level text as debug
	logTest := log.Ctx(ctx).With().Str("component", "test").Logger().Hook(logz.Hooks.DebugHook)
	// use info level for write but it will show debug
	logTest.Info().Msg("helloo level info but show debug")

	// use level with io.Copy
	_, _ = io.Copy(logz.LevelWriter(logTest, zerolog.DebugLevel), strings.NewReader("message X"))
}
