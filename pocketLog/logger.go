package pocketlog

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	output    io.Writer
	threshold Level
}

// New returns a pointer to a new Logger, which is then used to log messages from the application.
func New(output io.Writer, threshold Level) *Logger {
	return &Logger{
		output:    output,
		threshold: threshold,
	}
}

// Handles the actual logging of messages
func (l Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

// Log 'Debug' level messages.
func (l Logger) Debugf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold <= LevelDebug {
		l.logf(format+"\n", args...)
	}
}

// Log 'Info' level messages.
func (l Logger) Infof(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold <= LevelInfo {
		l.logf(format+"\n", args...)
	}
}

// Log 'Error' level messages.
func (l Logger) Errorf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold <= LevelError {
		l.logf(format+"\n", args...)
	}
}
