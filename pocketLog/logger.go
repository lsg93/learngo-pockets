package pocketlog

import "fmt"

type Logger struct {
	threshold Level
	shouldLog shouldLogRules
}

type shouldLogRules struct {
	Debug bool
	Info  bool
	Error bool
}

// New returns a pointer to a new Logger, which is then used to log messages from the application.
// We compute the levels at which messages should be logged based on the threshold passed when the struct is instantiated.
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
		shouldLog: shouldLogRules{
			Debug: LevelDebug >= threshold,
			Info:  LevelInfo >= threshold,
			Error: LevelError >= threshold,
		},
	}
}

// Log 'Debug' level messages.
func (l *Logger) Debugf(format string, a ...any) string {
	if !l.shouldLog.Debug {
		return ""
	}
	return fmt.Sprintf(format, a...)
}

// Log 'Info' level messages.
func (l *Logger) Infof(format string, a ...any) string {
	if !l.shouldLog.Info {
		return ""
	}
	return fmt.Sprintf(format, a...)
}

// Log 'Error' level messages.
func (l *Logger) Errorf(format string, a ...any) string {
	if !l.shouldLog.Error {
		return ""
	}
	return fmt.Sprintf(format, a...)
}
