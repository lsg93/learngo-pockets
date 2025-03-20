package pocketlog_own

import (
	"fmt"
	"io"
)

type logger struct {
	output    io.Writer
	threshold Level
	shouldLog shouldLogRules
}

type shouldLogRules map[Level]bool

// New returns a pointer to a new Logger, which is then used to log messages from the application.
// We compute the levels at which messages should be logged based on the threshold passed when the struct is instantiated.
func New(output io.Writer, threshold Level) *logger {
	return &logger{
		output:    output,
		threshold: threshold,
		shouldLog: shouldLogRules{
			LevelDebug: LevelDebug >= threshold,
			LevelInfo:  LevelInfo >= threshold,
			LevelError: LevelError >= threshold,
		},
	}
}

// Handles logging of functions to output
func (l *logger) logf(level Level, format string, args ...any) {
	shouldLog := l.shouldLog[level]
	if shouldLog {
		return
	}
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

// Log 'Debug' level messages.
func (l *logger) Debugf(format string, args ...any) {
	l.logf(LevelDebug, format+"\n", args...)
}

// Log 'Info' level messages.
func (l *logger) Infof(format string, args ...any) {
	l.logf(LevelInfo, format+"\n", args...)
}

// Log 'Error' level messages.
func (l *logger) Errorf(format string, args ...any) {
	l.logf(LevelError, format+"\n", args...)
}
