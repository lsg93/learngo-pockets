package pocketlog_own

import "io"

type LoggerOption func(*logger)

func WithOutput(output io.Writer) LoggerOption {
	return func(l *logger) {
		l.output = output
	}
}

func WithLogLevel() LoggerOption {
	return func(l *logger) {
		l.shouldLogLevel = true
	}
}

func WithThreshold(level LoggerLevel) LoggerOption {
	return func(l *logger) {
		l.threshold = level
	}
}
