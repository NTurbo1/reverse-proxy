package log

import (
	"log/slog"
	"os"
)

const (
	LevelTrace = slog.Level(-8)
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelFixMe = slog.Level(6)
	LevelError = slog.LevelError
	LevelFatal = slog.Level(12)
)

const (
	TRACE = "TRACE"
	DEBUG = "DEBUG"
	INFO = "INFO"
	WARN = "WARN"
	FIXME = "FIXME"
	ERROR = "ERROR"
	FATAL = "FATAL"
)

var logger *slog.Logger

func InitLogger(level string) {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: getLevel(level),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {

			// Customize the name of the level key and the output string, including
			// custom level values.
			if a.Key == slog.LevelKey {
				// Handle custom level values.
				level := a.Value.Any().(slog.Level)

				// This could also look up the name from a map or other structure, but
				// this demonstrates using a switch statement to rename levels. For
				// maximum performance, the string values should be constants, but this
				// example uses the raw strings for readability.
				switch {
				case level < LevelDebug:
					a.Value = slog.StringValue(TRACE)
				case level < LevelInfo:
					a.Value = slog.StringValue(DEBUG)
				case level < LevelWarn:
					a.Value = slog.StringValue(INFO)
				case level < LevelFixMe:
					a.Value = slog.StringValue(WARN)
				case level < LevelError:
					a.Value = slog.StringValue(FIXME)
				case level < LevelFatal:
					a.Value = slog.StringValue(ERROR)
				default:
					a.Value = slog.StringValue(FATAL)
				}
			}

			return a
		},
	}))
}

func getLevel(levelName string) slog.Level {
	switch levelName {
	case TRACE:
		return LevelTrace
	case DEBUG:
		return LevelDebug 
	case INFO:
		return LevelInfo 
	case WARN:
		return LevelWarn 
	case FIXME:
		return LevelFixMe
	case ERROR:
		return LevelError 
	case FATAL:
		return LevelFatal 
	default:
		return LevelInfo
	}
}
