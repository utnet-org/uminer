package log

import (
	"context"
	"fmt"
	commctx "uminer/common/context"
)

// Helper is a logger helper.
type Helper struct {
	logger Logger
}

// NewHelper new a logger helper.
func NewHelper(name string, logger Logger) *Helper {
	return &Helper{
		logger: logger,
	}
}

// Debug logs a message at debug level.
func (h *Helper) Debug(ctx context.Context, a ...interface{}) {
	h.logger.Printl(LevelDebug, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx), "message", fmt.Sprint(a...))
}

// Debugf logs a message at debug level.
func (h *Helper) Debugf(ctx context.Context, format string, a ...interface{}) {
	h.logger.Printl(LevelDebug, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx), "message", fmt.Sprintf(format, a...))
}

// Debugw logs a message at debug level.
func (h *Helper) Debugw(ctx context.Context, pairs ...interface{}) {
	h.logger.Printl(LevelDebug, append([]interface{}{commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx)}, pairs...)...)
}

// Info logs a message at info level.
func (h *Helper) Info(ctx context.Context, a ...interface{}) {
	h.logger.Printl(LevelInfo, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx), "message", fmt.Sprint(a...))
}

// Infof logs a message at info level.
func (h *Helper) Infof(ctx context.Context, format string, a ...interface{}) {
	h.logger.Printl(LevelInfo, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx), "message", fmt.Sprintf(format, a...))
}

// Infow logs a message at info level.
func (h *Helper) Infow(ctx context.Context, pairs ...interface{}) {
	h.logger.Printl(LevelInfo, append([]interface{}{commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx)}, pairs...)...)
}

// Warn logs a message at warn level.
func (h *Helper) Warn(ctx context.Context, a ...interface{}) {
	h.logger.Printl(LevelWarn, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx), "message", fmt.Sprint(a...))
}

// Warnf logs a message at warnf level.
func (h *Helper) Warnf(ctx context.Context, format string, a ...interface{}) {
	h.logger.Printl(LevelWarn, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx), "message", fmt.Sprintf(format, a...))
}

// Warnw logs a message at warnf level.
func (h *Helper) Warnw(ctx context.Context, pairs ...interface{}) {
	h.logger.Printl(LevelWarn, append([]interface{}{commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx)}, pairs...)...)
}

// Error logs a message at error level.
func (h *Helper) Error(ctx context.Context, a ...interface{}) {
	h.logger.Printl(LevelError, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx), "message", fmt.Sprint(a...))
}

// Errorf logs a message at error level.
func (h *Helper) Errorf(ctx context.Context, format string, a ...interface{}) {
	h.logger.Printl(LevelError, commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx), "message", fmt.Sprintf(format, a...))
}

// Errorw logs a message at error level.
func (h *Helper) Errorw(ctx context.Context, pairs ...interface{}) {
	h.logger.Printl(LevelError, append([]interface{}{commctx.RequestIdKey(), commctx.RequestIdFromContext(ctx)}, pairs...)...)
}
