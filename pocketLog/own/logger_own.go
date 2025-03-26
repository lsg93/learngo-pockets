package pocketlog_own

import (
	"fmt"
	"io"
	"os"
)

type logger struct {
	charLimit      int
	formatter      LoggerMessageFormatter
	output         io.Writer
	threshold      LoggerLevel
	shouldLog      shouldLogRules
	shouldLogLevel bool
}

type LoggerContext struct {
	level LoggerLevel
}

type shouldLogRules map[LoggerLevel]bool

// New returns a pointer to a new Logger, which is then used to log messages from the application.
func New(options ...LoggerOption) *logger {
	// The default value for the character limit of logged messages is 1000.
	// The default formatter is the PlaintextFormatter, which formats messages as they are given.
	// The default value for output is os.Stdout.
	// The default threshold at which error messages will be logged is LevelInfo.
	l := &logger{
		charLimit:      1000,
		formatter:      &PlaintextFormatter{},
		output:         os.Stdout,
		shouldLogLevel: false,
		threshold:      LevelInfo,
	}

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
	return l
}

// Handles logging of functions to output
func (l *logger) logf(ctx LoggerContext, format string, args ...any) {
	shouldLog := l.shouldLog[ctx.level]
	if !shouldLog {
		return
	}

	if l.shouldLogLevel {
		format = "[" + ctx.level.String() + "]" + format
	}

	if len(format) > l.charLimit {
		format = trimString(format, l.charLimit-3)
		format = format + "..."
	}

	format = l.formatter.Format(format, ctx)

	_, _ = fmt.Fprintf(l.output, format, args...)
}

// Log 'Debug' LoggerLevel messages.
func (l *logger) Debugf(format string, args ...any) {
	ctx := LoggerContext{level: LevelDebug}
	l.logf(ctx, format, args...)
}

// Log 'Info' LoggerLevel messages.
func (l *logger) Infof(format string, args ...any) {
	ctx := LoggerContext{level: LevelInfo}
	l.logf(ctx, format, args...)
}

// Log 'Error' LoggerLevel messages.
func (l *logger) Errorf(format string, args ...any) {
	ctx := LoggerContext{level: LevelError}
	l.logf(ctx, format, args...)
}

func trimString(s string, limit int) string {
	if len(s) <= limit {
		return s
	}

	// Find the rune boundary at the limit.
	runeCount := 0
	for i := range s {
		if runeCount >= limit {
			return s[:i]
		}
		runeCount++
	}

	return s // Should not reach here if limit is smaller than len(s)
}
