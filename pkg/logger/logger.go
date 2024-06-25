package logger

import (
	"context"
	"log/slog"
	"net"
	"os"
)

const (
	dev  = "dev"
	prod = "prod"
)

const (
	levelDebug = "LevelDebug"
	levelInfo  = "LevelInfo"
	levelWarn  = "LevelWarn"
	levelEvent = "LevelEvent"
	levelError = "LevelError"
)

const (
	eventLvl slog.Level = 5
)

type config interface {
	GetEnv() string
	GetLogAddr() string
	GetLevel() string
}

var LevelNames = map[slog.Leveler]string{
	eventLvl: "EVENT",
}

type Logger struct {
	sl *slog.Logger
}

func New(c config) (*Logger, error) {
	switch c.GetEnv() {
	case prod:
		conn, err := net.Dial("tcp", c.GetLogAddr())
		if err != nil {
			return nil, err
		}
		ss := slog.New(slog.NewJSONHandler(conn, &slog.HandlerOptions{
			AddSource:   false,
			Level:       parseLevel(c.GetLevel()),
			ReplaceAttr: replaceEventLevelName,
		}))

		return &Logger{sl: ss}, nil

	case dev:
		ss := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource:   false,
			Level:       parseLevel(c.GetLevel()),
			ReplaceAttr: replaceEventLevelName,
		}))

		return &Logger{sl: ss}, nil

	default:
		return &Logger{slog.New(slog.NewTextHandler(os.Stdout, nil))}, nil
	}
}

func parseLevel(lvl string) slog.Level {
	switch lvl {
	case levelDebug:
		return slog.LevelDebug

	case levelInfo:
		return slog.LevelInfo

	case levelWarn:
		return slog.LevelWarn

	case levelEvent:
		return eventLvl

	case levelError:
		return slog.LevelError

	default:
		return slog.LevelError
	}
}

func replaceEventLevelName(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		level := a.Value.Any().(slog.Level)
		levelLabel, exists := LevelNames[level]
		if !exists {
			levelLabel = level.String()
		}
		a.Value = slog.StringValue(levelLabel)
	}
	return a
}

func (l *Logger) Debug(msg string, args ...any) {
	l.sl.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.sl.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.sl.Warn(msg, args...)
}

func (l *Logger) Event(msg string, args ...any) {
	l.sl.Log(context.Background(), eventLvl, msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.sl.Error(msg, args...)
}
