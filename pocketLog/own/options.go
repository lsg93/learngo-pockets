package pocketlog_own

import "io"

type LoggerOption func(*logger)

func WithCharacterLimit(limit int) LoggerOption {
	return func(l *logger) {
		if limit <= 0 {
			return
		}
		l.charLimit = limit
	}
}

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
