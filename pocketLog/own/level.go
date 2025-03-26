package pocketlog_own

import "fmt"

type LoggerLevel byte

// Enum of thresholds at which messages should be logged.
const (
	LevelDebug LoggerLevel = iota
	LevelInfo
	LevelError
)

// String representation of enum values - fmt.Stringer
func (l LoggerLevel) String() string {
	switch l {
	case LevelDebug:
		return "Debug"
	case LevelInfo:
		return "Info"
	case LevelError:
		return "Error"
	default:
		return fmt.Sprintf("%d", int(l))
	}
}
