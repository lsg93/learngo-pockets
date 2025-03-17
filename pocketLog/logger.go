package pocketlog

import "fmt"

type Logger struct {
	threshold Level
}

// New returns a pointer to a new Logger, which is then used to log messages from the application.
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

// Log 'Debug' level messages.
func (l *Logger) Debugf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

// Log 'Info' level messages.
func (l *Logger) Infof(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

// Log 'Error' level messages.
func (l *Logger) Errorf(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}
