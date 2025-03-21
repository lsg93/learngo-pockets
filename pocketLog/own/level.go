package pocketlog_own

type LoggerLevel byte

// Enum of thresholds at which messages should be logged.
const (
	LevelDebug LoggerLevel = iota
	LevelInfo
	LevelError
)
