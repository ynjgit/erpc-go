package log

import (
	"fmt"
	std "log"
)

func init() {
	flags := std.Flags()
	flags = flags | std.Lshortfile
	std.SetFlags(flags)
}

// Logger the logging interface
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Sync() error

	WithFields(fields ...string) Logger
}

var defaultLogger Logger
var stdLogger = &stdLoggerWrapper{callerSkip: 3}

func setLogger(l Logger) {
	defaultLogger = l
}

func getLogger() Logger {
	if defaultLogger != nil {
		return defaultLogger
	}

	return stdLogger
}

type stdLoggerWrapper struct {
	callerSkip int
}

// Debug ...
func (l *stdLoggerWrapper) Debug(args ...interface{}) {
	std.Output(l.callerSkip, fmt.Sprint(args...))
}

// Debugf ...
func (l *stdLoggerWrapper) Debugf(format string, args ...interface{}) {
	std.Output(l.callerSkip, fmt.Sprintf(format, args...))
}

// Info ...
func (l *stdLoggerWrapper) Info(args ...interface{}) {
	std.Output(l.callerSkip, fmt.Sprint(args...))
}

// Infof ...
func (l *stdLoggerWrapper) Infof(format string, args ...interface{}) {
	std.Output(l.callerSkip, fmt.Sprintf(format, args...))
}

// Warn ...
func (l *stdLoggerWrapper) Warn(args ...interface{}) {
	std.Output(l.callerSkip, fmt.Sprint(args...))
}

// Warnf ...
func (l *stdLoggerWrapper) Warnf(format string, args ...interface{}) {
	std.Output(l.callerSkip, fmt.Sprintf(format, args...))

}

// Error ...
func (l *stdLoggerWrapper) Error(args ...interface{}) {
	std.Output(l.callerSkip, fmt.Sprint(args...))
}

// Errorf ...
func (l *stdLoggerWrapper) Errorf(format string, args ...interface{}) {
	std.Output(l.callerSkip, fmt.Sprintf(format, args...))
}

// Fatal ...
func (l *stdLoggerWrapper) Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// Fatalf ...
func (l *stdLoggerWrapper) Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

// WithFields ...
func (l *stdLoggerWrapper) WithFields(fields ...string) Logger {
	return l
}

// Sync ...
func (l *stdLoggerWrapper) Sync() error {
	return nil
}
