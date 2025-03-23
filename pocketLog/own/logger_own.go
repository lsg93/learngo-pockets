package pocketlog_own

import (
	"fmt"
	"io"
	"os"
)

type logger struct {
	logLevels      logLevels
	output         io.Writer
	threshold      LoggerLevel
	shouldLog      shouldLogRules
	shouldLogLevel bool
}

type shouldLogRules map[LoggerLevel]bool

type logLevels map[LoggerLevel]string

// New returns a pointer to a new Logger, which is then used to log messages from the application.

func New(options ...LoggerOption) *logger {
	// The default value for output is os.Stdout.
	// The default threshold at which error messages will be logged is LevelInfo.
	l := &logger{output: os.Stdout, shouldLogLevel: false, threshold: LevelInfo}

	for _, option := range options {
		option(l)
	}

	// We compute the levels at which messages should be logged based on the threshold passed when the struct is instantiated.
	l.shouldLog = shouldLogRules{
		LevelDebug: LevelDebug >= l.threshold,
		LevelInfo:  LevelInfo >= l.threshold,
		LevelError: LevelError >= l.threshold,
	}

	// If the flag showLogLevel is true, then we print the log level with the provided message.
	l.logLevels = logLevels{
		LevelDebug: "[DEBUG]: ",
		LevelInfo:  "[INFO]: ",
		LevelError: "[ERROR]: ",
	}

	return l
}

// Handles logging of functions to output
func (l *logger) logf(LoggerLevel LoggerLevel, format string, args ...any) {
	shouldLog := l.shouldLog[LoggerLevel]
	if !shouldLog {
		return
	}

	if l.shouldLogLevel {
		format = l.logLevels[LoggerLevel] + format
	}

	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

// Log 'Debug' LoggerLevel messages.
func (l *logger) Debugf(format string, args ...any) {
	l.logf(LevelDebug, format+"\n", args...)
}

// Log 'Info' LoggerLevel messages.
func (l *logger) Infof(format string, args ...any) {
	l.logf(LevelInfo, format+"\n", args...)
}

// Log 'Error' LoggerLevel messages.
func (l *logger) Errorf(format string, args ...any) {
	l.logf(LevelError, format+"\n", args...)
}
