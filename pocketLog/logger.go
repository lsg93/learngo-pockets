package pocketlog

import "fmt"

type Logger struct {
	threshold Level
}

func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

func (l *Logger) Debugf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func (l *Logger) Infof(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func (l *Logger) Errorf(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}
