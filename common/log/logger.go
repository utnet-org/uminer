package log

import (
	"context"
	"strings"

	klog "github.com/go-kratos/kratos/v2/log"
)

type Level int32

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelSilent
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelSilent:
		return "SILENT"
	}
	return ""
}

func ConvertFromString(levelStr string) Level {
	switch strings.ToUpper(levelStr) {
	case LevelDebug.String():
		return LevelDebug
	case LevelInfo.String():
		return LevelInfo
	case LevelWarn.String():
		return LevelWarn
	case LevelError.String():
		return LevelError
	case LevelSilent.String():
		return LevelSilent
	}
	return LevelInfo
}

type Logger interface {
	Log(level klog.Level, pairs ...interface{}) error
	Printl(level Level, pairs ...interface{})
	Print(pairs ...interface{})
	ResetLevel(level Level)
	Level() Level
}

var (
	// 默认logger
	DefaultLogger Logger = NewStdLogger()
	// 默认helper
	defaultHelper = NewHelper("", DefaultLogger)
)

// Debug logs a message at debug level.
func Debug(ctx context.Context, a ...interface{}) {
	defaultHelper.Debug(ctx, a...)
}

// Debugf logs a message at debug level.
func Debugf(ctx context.Context, format string, a ...interface{}) {
	defaultHelper.Debugf(ctx, format, a...)
}

// Debugw logs a message at debug level.
func Debugw(ctx context.Context, pairs ...interface{}) {
	defaultHelper.Debugw(ctx, pairs...)
}

// Info logs a message at info level.
func Info(ctx context.Context, a ...interface{}) {
	defaultHelper.Info(ctx, a...)
}

// Infof logs a message at info level.
func Infof(ctx context.Context, format string, a ...interface{}) {
	defaultHelper.Infof(ctx, format, a...)
}

// Infow logs a message at info level.
func Infow(ctx context.Context, pairs ...interface{}) {
	defaultHelper.Infow(ctx, pairs...)
}

// Warn logs a message at warn level.
func Warn(ctx context.Context, a ...interface{}) {
	defaultHelper.Warn(ctx, a...)
}

// Warnf logs a message at warnf level.
func Warnf(ctx context.Context, format string, a ...interface{}) {
	defaultHelper.Warnf(ctx, format, a...)
}

// Warnw logs a message at warnf level.
func Warnw(ctx context.Context, pairs ...interface{}) {
	defaultHelper.Warnw(ctx, pairs...)
}

// Error logs a message at error level.
func Error(ctx context.Context, a ...interface{}) {
	defaultHelper.Error(ctx, a...)
}

// Errorf logs a message at error level.
func Errorf(ctx context.Context, format string, a ...interface{}) {
	defaultHelper.Errorf(ctx, format, a...)
}

// Errorw logs a message at error level.
func Errorw(ctx context.Context, pairs ...interface{}) {
	defaultHelper.Errorw(ctx, pairs...)
}
