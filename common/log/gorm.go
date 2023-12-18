package log

import (
	"context"
	"errors"
	"fmt"
	"time"
	commctx "uminer/common/context"

	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

var (
	DefaultGormLogger = NewGormLogger(Config{
		SlowThreshold: 200 * time.Millisecond,
	}, NewStdLogger())
)

type gormLogger struct {
	logger                    Logger
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
}

func NewGormLogger(config Config, logger Logger) glogger.Interface {
	l := &gormLogger{
		logger:                    logger,
		SlowThreshold:             config.SlowThreshold,
		IgnoreRecordNotFoundError: config.IgnoreRecordNotFoundError,
	}
	return l
}

func ConvertToGorm(level Level) glogger.LogLevel {
	switch level {
	case LevelSilent:
		return glogger.Silent
	case LevelError:
		return glogger.Error
	case LevelWarn:
		return glogger.Warn
	case LevelInfo:
		return glogger.Info
	}

	return glogger.Info
}

func convertFromGorm(level glogger.LogLevel) Level {
	switch level {
	case glogger.Silent:
		return LevelSilent
	case glogger.Error:
		return LevelError
	case glogger.Warn:
		return LevelWarn
	case glogger.Info:
		return LevelInfo
	}

	return LevelInfo
}

type Config struct {
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
}

func (l *gormLogger) LogMode(level glogger.LogLevel) glogger.Interface {
	l.logger.ResetLevel(convertFromGorm(level))
	return l
}

func (l *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Printl(LevelInfo, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx), "message", fmt.Sprintf(msg, data...))

}

// Warn print warn messages
func (l *gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Printl(LevelWarn, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx), "message", fmt.Sprintf(msg, data...))
}

// Error print error messages
func (l *gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Printl(LevelError, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx), "message", fmt.Sprintf(msg, data...))
}

const (
	traceStr     = "[%.3fms] [rows:%v] %s"
	traceWarnStr = "%s\n[%.3fms] [rows:%v] %s"
	traceErrStr  = "%s\n[%.3fms] [rows:%v] %s"
)

// Trace print sql message
func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.logger.Level() < LevelSilent {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.logger.Level() <= LevelError && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
			sql, rows := fc()
			if rows == -1 {
				l.logger.Printl(LevelError, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx),
					"message", fmt.Sprintf(traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, "-", sql))
			} else {
				l.logger.Printl(LevelError, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx),
					"message", fmt.Sprintf(traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, rows, sql))
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.logger.Level() <= LevelWarn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			if rows == -1 {
				l.logger.Printl(LevelWarn, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx),
					"message", fmt.Sprintf(traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql))
			} else {
				l.logger.Printl(LevelWarn, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx),
					"message", fmt.Sprintf(traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql))
			}
		case l.logger.Level() == LevelInfo:
			sql, rows := fc()
			if rows == -1 {
				l.logger.Printl(LevelInfo, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx),
					"message", fmt.Sprintf(traceStr, float64(elapsed.Nanoseconds())/1e6, "-", sql))
			} else {
				l.logger.Printl(LevelInfo, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx),
					"message", fmt.Sprintf(traceStr, float64(elapsed.Nanoseconds())/1e6, rows, sql))
			}
		}
	}
}
