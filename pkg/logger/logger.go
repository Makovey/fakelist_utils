package logger

type Logger interface {
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Infof(format string, args ...any)
	Warn(format string, args ...any)
	Warnf(format string, args ...any)
	Error(format string, args ...any)
	Errorf(format string, args ...any)
}
