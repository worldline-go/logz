package logz

import (
	"os"
	"strconv"

	"github.com/mattn/go-isatty"
)

type Selection uint8

const (
	SelectAuto Selection = iota
	SelectTrue
	SelectFalse
)

var Default = struct {
	TimeStamp bool
	Caller    bool
	Pretty    Selection
	Level     string
}{
	TimeStamp: true,
	Caller:    true,
	Pretty:    SelectAuto,
	Level:     "info",
}

func checkDefault(p *bool, v bool) bool {
	if p == nil {
		return v
	}

	return *p
}

func checkPretty(p *bool, v Selection) bool {
	if p == nil {
		switch v {
		case SelectAuto:
			v, ok := os.LookupEnv(EnvPretty)
			if ok {
				result, _ := strconv.ParseBool(v)

				return result
			}

			return isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
		case SelectFalse:
			return false
		case SelectTrue:
			return true
		}
	}

	return *p
}

func checkLevel(p *string, level string) string {
	if p != nil {
		return *p
	}

	if v, ok := os.LookupEnv(EnvLevel); ok {
		return v
	}

	return level
}
