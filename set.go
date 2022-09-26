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

type Selection uint8

const (
	Auto Selection = iota
	True
	False
)

type Settings struct {
	Pretty    Selection
	Timestamp bool
	Caller    bool
	_isPretty bool
}

func (s Settings) SetTimestamp(v bool) *Settings {
	s.Timestamp = v

	return &s
}

func (s Settings) SetCaller(v bool) *Settings {
	s.Caller = v

	return &s
}

func (s Settings) SetPretty(v Selection) *Settings {
	s.Pretty = v

	return &s
}

var DefaultLogSettings = Settings{
	Pretty:    Auto,
	Timestamp: true,
	Caller:    false,
}

// InitializeLog choice between json format or common format.
// LOG_PRETTY boolean environment value always override the decision.
// Override with 'pretty' boolean argument.
func InitializeLog(lSettings *Settings) {
	logger := GetLogger(lSettings)
	log.Logger = logger

	zerolog.DefaultContextLogger = &log.Logger
}

func GetLogger(lSettings *Settings) (logger zerolog.Logger) {
	settings := DefaultLogSettings
	if lSettings != nil {
		settings = *lSettings
	}

	defer func() {
		zerolog.TimeFieldFormat = TimeFormat

		var logX zerolog.Context
		if settings._isPretty {
			logX = zerolog.New(LogWriter).With()
		} else {
			logX = zerolog.New(os.Stderr).With()
		}

		if settings.Timestamp {
			logX = logX.Timestamp()
		}

		if settings.Caller {
			logX = logX.Caller()
		}

		logger = logX.Logger()
	}()

	// set pretty format
	if settings.Pretty != Auto {
		settings._isPretty = settings.Pretty == True

		return
	}

	if v, ok := os.LookupEnv(PrettyEnv); ok {
		settings._isPretty, _ = strconv.ParseBool(v)

		return
	}

	settings._isPretty = isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())

	return
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
