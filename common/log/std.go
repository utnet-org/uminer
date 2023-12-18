package log

import (
	"os"
	"sync/atomic"

	klog "github.com/go-kratos/kratos/v2/log"
)

type stdLogger struct {
	logger klog.Logger
	level  Level
}

func NewStdLogger() Logger {
	l := &stdLogger{
		//logger: klog.NewStdLogger(log.Writer()),
		logger: klog.With(klog.NewStdLogger(os.Stdout),
			"ts", klog.DefaultTimestamp,
		),
	}
	l.setLevel(LevelInfo)

	return l
}

func (l *stdLogger) Log(level klog.Level, pairs ...interface{}) error {
	if level < klog.Level(l.getLevel()) {
		return nil
	}

	p := make([]interface{}, 0)
	p = append(p, []interface{}{"level", level}...)
	p = append(p, pairs...)
	l.logger.Log(klog.Level(level), p...)
	return nil
}

func (l *stdLogger) Printl(level Level, pairs ...interface{}) {
	if level < l.getLevel() {
		return
	}

	p := make([]interface{}, 0)
	p = append(p, []interface{}{"level", level}...)
	p = append(p, pairs...)
	l.logger.Log(klog.Level(level), p...)
}

func (l *stdLogger) Print(pairs ...interface{}) {
	l.logger.Log(klog.Level(l.getLevel()), pairs...)
}

func (l *stdLogger) ResetLevel(level Level) {
	l.setLevel(level)
}

func (l *stdLogger) Level() Level {
	return l.getLevel()
}

func (l *stdLogger) getLevel() Level {
	return Level(atomic.LoadInt32((*int32)(&l.level)))
}

func (l *stdLogger) setLevel(level Level) {
	atomic.StoreInt32((*int32)(&l.level), int32(level))
}
