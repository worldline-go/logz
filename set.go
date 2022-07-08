package logz

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	PrettyEnv        = "LOG_PRETTY"
	TimeFormat       = time.RFC3339Nano
	TimePrettyFormat = "2006-01-02 15:04:05"
)

var LogWriter = zerolog.ConsoleWriter{
	Out: os.Stderr,
	FormatTimestamp: func(i interface{}) string {
		parse, _ := time.Parse(TimeFormat, i.(string))

		return parse.Format(TimePrettyFormat)
	},
}

// InitializeLog choice between json format or common format.
// LOG_PRETTY boolean environment value always override the decision.
// Override with 'pretty' boolean argument.
func InitializeLog(pretty *bool) {
	isPretty := false

	defer func() {
		zerolog.TimeFieldFormat = TimeFormat
		if isPretty {
			log.Logger = zerolog.New(LogWriter).With().Timestamp().Logger()
		}

		zerolog.DefaultContextLogger = &log.Logger
	}()

	if pretty != nil {
		isPretty = *pretty

		return
	}

	if v, ok := os.LookupEnv(PrettyEnv); ok {
		isPretty, _ = strconv.ParseBool(v)

		return
	}

	isPretty = isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
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
